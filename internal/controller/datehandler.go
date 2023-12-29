package controller

import (
	"fmt"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/m00n-arch/remind-stuff-bot/internal/db"
	"github.com/m00n-arch/remind-stuff-bot/internal/languages"
)

func (c *Controller) DateHandler(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, languages.MessageEnterText)

	parse, err := time.Parse("02.01.2006 15:04", update.Message.Text)
	if err != nil {
		return fmt.Errorf("wrong time parsing format, try again: %w", err)
	}
	err = c.alertDB.AddAlert(db.Alert{
		UserID: strconv.FormatInt(update.Message.From.ID, 10),
		Date:   parse,
	})
	if err != nil {
		return err
	}

	err = c.userDB.UpdateState(strconv.FormatInt(update.Message.Chat.ID, 10), db.CreateTextState)
	if err != nil {
		return err
	}

	if _, err := c.bot.Send(msg); err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}
	return nil
}
