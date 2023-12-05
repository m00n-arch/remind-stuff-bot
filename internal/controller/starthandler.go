package controller

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m00n-arch/remind-stuff-bot/internal/buttons"
)

func (c *Controller) StartHandler(update tgbotapi.Update) error {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите интересующее вас действие")
	msg.ReplyMarkup = buttons.MainMenuButtons

	if _, err := c.bot.Send(msg); err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}
	return nil
}
