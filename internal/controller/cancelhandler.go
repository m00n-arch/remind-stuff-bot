package controller

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m00n-arch/remind-stuff-bot/internal/buttons"
	"github.com/m00n-arch/remind-stuff-bot/internal/db"
)

func (c *Controller) CancelHandler(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите интересующее вас действие")

	err := c.userDB.UpdateState(strconv.FormatInt(update.Message.Chat.ID, 10), db.StartState)
	if err != nil {
		return err
	}
	msg.Text = "Выберите интересующее вас действие"
	msg.ReplyMarkup = buttons.MainMenuButtons

	return nil
}
