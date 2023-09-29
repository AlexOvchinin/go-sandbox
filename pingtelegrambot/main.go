package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	tele "gopkg.in/telebot.v3"
)

type Chat struct {
	ID        int64
	usernames mapset.Set[string]
	users     map[int64]string
}

var chats = make(map[int64]*Chat)

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

	b.Handle("/add", handleAddCommand)
	b.Handle("/everyone", handleEveryoneCommand)
	b.Handle(tele.OnUserJoined, handleUserJoined)
	b.Handle(tele.OnUserLeft, handleUserLeft)

	b.Start()
}

func handleAddCommand(ctx tele.Context) error {
	chat := chats[ctx.Chat().ID]
	if chat == nil {
		chat = &Chat{ctx.Chat().ID, mapset.NewSet[string](), make(map[int64]string)}
		chats[ctx.Chat().ID] = chat
	}
	for _, entity := range ctx.Message().Entities {
		switch entity.Type {
		case tele.EntityMention:
			chat.usernames.Add(ctx.Message().EntityText(entity))
		case tele.EntityTMention:
			fmt.Println(entity.User)
			chat.users[entity.User.ID] = entity.User.FirstName
			chat.usernames.Remove(entity.User.Username)
		}
	}
	return ctx.Send("Added")
}

func handleEveryoneCommand(ctx tele.Context) error {
	chat := chats[ctx.Chat().ID]
	if chat == nil {
		return ctx.Send("Noone to mention. Please use /add to add users to mention manually")
	}
	var builder strings.Builder
	usernamesIter := chat.usernames.Iterator()
	for elem := range usernamesIter.C {
		fmt.Fprintf(&builder, "%v", elem)
		fmt.Fprintf(&builder, " ")
	}
	for userId, username := range chat.users {
		fmt.Fprintf(&builder, "[%v](tg://user?id=%v)", username, userId)
		fmt.Fprintf(&builder, " ")
	}
	return ctx.Send(builder.String(), tele.ModeMarkdownV2)
}

func handleUserJoined(ctx tele.Context) error {
	chat := chats[ctx.Chat().ID]
	if chat == nil {
		chat = &Chat{ctx.Chat().ID, mapset.NewSet[string](), make(map[int64]string)}
		chats[ctx.Chat().ID] = chat
	}
	joinedUser := ctx.Message().UserJoined
	chat.users[joinedUser.ID] = joinedUser.FirstName
	chat.usernames.Remove(joinedUser.Username)
	return nil
}

func handleUserLeft(ctx tele.Context) error {
	chat := chats[ctx.Chat().ID]
	if chat == nil {
		chat = &Chat{ctx.Chat().ID, mapset.NewSet[string](), make(map[int64]string)}
		chats[ctx.Chat().ID] = chat
	}
	leftUser := ctx.Message().UserLeft
	delete(chat.users, leftUser.ID)
	chat.usernames.Remove(leftUser.Username)
	return nil
}
