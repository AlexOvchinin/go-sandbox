package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	tele "gopkg.in/telebot.v3"
)

type Chat struct {
	ID       int64
	mentions mapset.Set[string]
	users    mapset.Set[int64]
}

var chats = make(map[int64]*Chat)

func main() {
	token := os.Getenv("PING_BOT_TOKEN")
	pref := tele.Settings{
		Token:   token,
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		Verbose: true,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/add", handleAdd)
	b.Handle(tele.OnAddedToGroup, handleEvent)
	b.Handle(tele.OnAddedToGroup, handleEvent)
	b.Handle(tele.OnUserJoined, handleEvent)
	b.Handle(tele.OnUserLeft, handleEvent)

	b.Start()
}

func handleAdd(c tele.Context) error {
	chatId := c.Chat().ID
	chat := chats[c.Chat().ID]
	if chat == nil {
		chat = &Chat{chatId, mapset.NewSet[string](), mapset.NewSet[int64]()}
		chats[c.Chat().ID] = chat
	}
	return c.Send(fmt.Sprintf("Hello %v!", chat))
	// for _, entity := range c.Message().Entities {
	// 	switch entity.Type {
	// 	case tele.EntityMention:
	// 	case tele.EntityTMention:
	// 	}
	// 	if entity.Type = tele.EntityMention {

	// 	}
	// 	if (entity.Type = EnttityType.)
	// 	fmt.Println(entity)
	// }
}

func handleEveryoneCommand(c tele.Context) error {
	return nil
}

func handUserJoined(c tele.Context) error {
	return nil
}

func handleUserLeft(c tele.Context) error {
	return nil
}

func handleEvent(c tele.Context) error {
	log.Println(c)
	message := fmt.Sprintf("%#v", c)

	log.Println(c.Chat())
	return c.Send(message)
}
