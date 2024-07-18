package main

import (
	"encoding/json"
	"fm/pingtelegrambot/handlers"
	"fm/pingtelegrambot/model"
	"fmt"
	"log"
	"os"
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

var storage = model.NewChatStorage(dataPath)

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

	handlers.Storage = storage

	loadChats()

	// bot commands
	b.Handle("/add", handlers.HandleAddCommand)
	b.Handle("/everyone", handlers.HandleEveryoneCommand)

	// chat events
	b.Handle(tele.OnUserJoined, handlers.HandleUserJoined)
	b.Handle(tele.OnUserLeft, handlers.HandleUserLeft)
	b.Handle(tele.OnMigration, handlers.HandleMigration)

	b.Start()

	log.Println("Ping Telegram Bot Started")
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

	fmt.Printf("Loaded %v chats\n", len(chats))

	for _, chat := range chats {
		users := []*model.User{}
		for id, firstName := range chat.Users {
			users = append(users, &model.User{
				ID:        id,
				FirstName: firstName,
			})
		}
		for username := range chat.Usernames {
			if len(username) > 0 {
				users = append(users, &model.User{
					Username: username,
				})
			}
		}
		storage.AddUsersToMention(chat.ID, model.MentionEveryoneName, users)
	}

	fmt.Println("Transferred old chats to the new model")
}
