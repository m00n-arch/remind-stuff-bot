package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m00n-arch/remind-stuff-bot/internal/buttons"
	"github.com/m00n-arch/remind-stuff-bot/internal/languages"
)

func (c *Controller) NewRemHandler(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите интересующее вас действие")

	msg.ReplyMarkup = buttons.Cancel
	msg.Text = languages.DateTimeFormat
}
