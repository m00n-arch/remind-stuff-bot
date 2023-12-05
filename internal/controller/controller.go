package controller

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m00n-arch/remind-stuff-bot/internal/db"
	"github.com/m00n-arch/remind-stuff-bot/internal/languages"
)

type Controller struct {
	userDB  *db.UsersDB
	alertDB *db.AlertsDB
	bot     *tgbotapi.BotAPI
}

func NewController(userDB *db.UsersDB, alertDB *db.AlertsDB, bot *tgbotapi.BotAPI) *Controller {
	return &Controller{
		bot:     bot,
		userDB:  userDB,
		alertDB: alertDB,
	}
}

func (c *Controller) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {

		userState, err := c.userDB.GetState(strconv.Itoa(int(update.Message.Chat.ID)))
		if err != nil {
			// handle later
			continue
		}

		// Extract the command from the Message.
		switch {
		case update.Message.Text == "/start":
			err = c.StartHandler(update)
		case update.Message.Text == languages.NewReminderButton:
			err = c.NewRemHandler(update)
		case update.Message.Text == languages.CancelButton:
			err = c.CancelHandler(update)
		case userState == db.CreateState:
			err = c.DateHandler(update)
		case userState == db.CreateTextState:
			err = c.TextHandler(update)
		default:
			err = c.DefaultHandler(update)
		}
		if err != nil {
			return err
		}
	}
	return fmt.Errorf("unreachable")
}
