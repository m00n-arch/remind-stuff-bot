package controller

import (
	"fmt"
	"log"
	"strconv"

	"ReminderBot/internal/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Controller struct {
	userDB *db.UsersDB
	bot    *tgbotapi.BotAPI
}

func NewController(userDB *db.UsersDB, bot *tgbotapi.BotAPI) *Controller {
	return &Controller{
		bot:    bot,
		userDB: userDB,
	}
}

func (c *Controller) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "ру ру")

		userState, err := c.userDB.GetState(strconv.Itoa(int(update.Message.Chat.ID)))
		if err != nil {
			// todo: handle later
			continue
		}

		// Extract the command from the Message.
		switch {
		case update.Message.Text == "/help":
			msg.Text = "Привет! Я бот-напоминатель. Используй команду /setreminder, чтобы установить напоминание."
		case update.Message.Text == "/create":
			err := c.userDB.UpdateState(strconv.FormatInt(update.Message.Chat.ID, 10), "createState")
			if err != nil {
				return err
			}
			msg.Text = "Введите дату и время для напоминания"
		case userState == "createState":
			// date handler
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := c.bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

	return fmt.Errorf("unreachable")
}
