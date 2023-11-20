package main

import (
	"log"
	"os"

	"ReminderBot/internal/controller"
	"ReminderBot/internal/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		return err
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	udb, err := db.NewUserDB("user.csv")
	if err != nil {
		return err
	}

	c := controller.NewController(udb, bot)

	return c.Run()
}
