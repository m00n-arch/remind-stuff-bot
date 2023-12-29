package controller

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m00n-arch/remind-stuff-bot/internal/languages"
)

func (c *Controller) ErrorHandler(update tgbotapi.Update, err error) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, languages.MessageError)

	if _, err := c.bot.Send(msg); err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}
	return nil
}
