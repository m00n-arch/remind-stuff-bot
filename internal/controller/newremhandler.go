package controller

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m00n-arch/remind-stuff-bot/internal/buttons"
	"github.com/m00n-arch/remind-stuff-bot/internal/db"
	"github.com/m00n-arch/remind-stuff-bot/internal/languages"
)

func (c *Controller) NewRemHandler(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите интересующее вас действие")

	err := c.userDB.UpdateState(strconv.FormatInt(update.Message.Chat.ID, 10), db.CreateState)
	if err != nil {
		return err
	}

	msg.ReplyMarkup = buttons.Cancel
	msg.Text = languages.MessageEnterDateTime

	if _, err := c.bot.Send(msg); err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}
	return nil
}
