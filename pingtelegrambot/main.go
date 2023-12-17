package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	tele "gopkg.in/telebot.v3"
)

var Empty struct{}

type Chat struct {
	ID        int64
	Usernames map[string]struct{}
	Users     map[int64]string
}

var chats = make(map[int64]*Chat)

var dataPath = os.Getenv("PING_BOT_DATA_PATH")
var mu sync.Mutex

func main() {
	token := os.Getenv("PING_BOT_TOKEN")
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	loadChats()

	b.Handle("/add", handleAddCommand)
	b.Handle("/everyone", handleEveryoneCommand)
	b.Handle(tele.OnUserJoined, handleUserJoined)
	b.Handle(tele.OnUserLeft, handleUserLeft)
	b.Handle(tele.OnMigration, handleMigration)

	b.Start()

	log.Println("Ping Telegram Bot Started")
}

func saveChats() {
	mu.Lock()
	defer mu.Unlock()

	f, err := os.Create(dataPath)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	marshaledChats, err := json.Marshal(chats)
	if err != nil {
		log.Println(err)
	}

	_, err = f.WriteString(string(marshaledChats))
	if err != nil {
		log.Println(err)
	}
}

func loadChats() {
	marshaledChats, err := os.ReadFile(dataPath)
	if err != nil {
		log.Println(err)
		fmt.Printf("No chats to load")
		return
	}

	err = json.Unmarshal(marshaledChats, &chats)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded %v chats", len(chats))
}

func handleAddCommand(ctx tele.Context) error {
	chat := getChat(ctx.Chat().ID)
	for _, entity := range ctx.Message().Entities {
		switch entity.Type {
		case tele.EntityMention:
			chat.Usernames[ctx.Message().EntityText(entity)] = Empty
		case tele.EntityTMention:
			chat.Users[entity.User.ID] = entity.User.FirstName
			delete(chat.Usernames, entity.User.Username)
		}
	}
	go saveChats()
	return ctx.Send("Added")
}

func handleEveryoneCommand(ctx tele.Context) error {
	chat := getChat(ctx.Chat().ID)

	var builder strings.Builder

	senderUsername := ctx.Message().Sender.Username
	for username := range chat.Usernames {
		if senderUsername != username {
			fmt.Fprintf(&builder, "%v", username)
			fmt.Fprintf(&builder, " ")
		}
	}

	senderId := ctx.Message().Sender.ID
	for userId, username := range chat.Users {
		if userId != senderId {
			fmt.Fprintf(&builder, "[%v](tg://user?id=%v)", username, userId)
			fmt.Fprintf(&builder, " ")
		}
	}
	message := builder.String()
	if len(message) == 0 {
		return ctx.Send("Noone to mention. Please use /add to add users to mention manually")
	}

	return ctx.Send(builder.String(), tele.ModeMarkdownV2)
}

func handleUserJoined(ctx tele.Context) error {
	chat := getChat(ctx.Chat().ID)
	joinedUser := ctx.Message().UserJoined
	chat.Users[joinedUser.ID] = joinedUser.FirstName
	delete(chat.Usernames, joinedUser.Username)
	go saveChats()
	return nil
}

func handleUserLeft(ctx tele.Context) error {
	chat := getChat(ctx.Chat().ID)
	leftUser := ctx.Message().UserLeft
	delete(chat.Users, leftUser.ID)
	go saveChats()
	return nil
}

func handleMigration(ctx tele.Context) error {
	migrageFrom := ctx.Message().MigrateFrom
	migrateTo := ctx.Message().MigrateTo
	chat := chats[migrageFrom]
	if chat != nil {
		chats[migrateTo] = chat
		delete(chats, migrageFrom)
	}
	return nil
}

func getChat(id int64) *Chat {
	chat := chats[id]
	if chat == nil {
		chat = &Chat{id, make(map[string]struct{}), make(map[int64]string)}
		chats[id] = chat
	}
	return chat
}
