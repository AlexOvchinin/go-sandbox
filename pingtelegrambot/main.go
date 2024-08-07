package main

import (
	"fm/pingtelegrambot/handlers"
	"fm/pingtelegrambot/model"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

var storage *model.ChatStorage

func main() {
	storage = model.NewChatStorage(os.Getenv("PING_BOT_DATA_PATH"))

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
