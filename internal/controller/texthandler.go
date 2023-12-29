package controller

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/m00n-arch/remind-stuff-bot/internal/db"
	"github.com/m00n-arch/remind-stuff-bot/internal/languages"
)

func (c *Controller) TextHandler(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, languages.MessageCreationFinished)

	err := c.alertDB.UpdateAlert(db.Alert{
		UserID:  strconv.FormatInt(update.Message.From.ID, 10),
		Content: update.Message.Text,
	})
	if err != nil {
		return err
	}

	err = c.userDB.UpdateState(strconv.FormatInt(update.Message.Chat.ID, 10), db.StartState)
	if err != nil {
		return err
	}

	if _, err := c.bot.Send(msg); err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}
	return nil
}
