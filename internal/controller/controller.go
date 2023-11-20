package controller

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m00n-arch/remind-stuff-bot/internal/buttons"
	"github.com/m00n-arch/remind-stuff-bot/internal/db"
	"github.com/m00n-arch/remind-stuff-bot/internal/languages"
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
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "te-hee")

		userState, err := c.userDB.GetState(strconv.Itoa(int(update.Message.Chat.ID)))
		if err != nil {
			// todo: handle later
			continue
		}

		// Extract the command from the Message.
		switch {
		case update.Message.Text == "/start":
			msg.ReplyMarkup = buttons.MainMenuButtons
			msg.Text = "Выберите интересующее вас действие"
		case update.Message.Text == languages.NewReminderButton:
			msg.ReplyMarkup = buttons.Cancel
			msg.Text = languages.DateTimeFormat
		case update.Message.Text == languages.CancelButton
			// handle later
		case update.Message.Text == "/help":
			msg.Text = "Привет! Я бот-напоминатель."
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
