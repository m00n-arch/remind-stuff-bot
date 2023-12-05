package controller

import (
	"fmt"
	"regexp"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m00n-arch/remind-stuff-bot/internal/db"
	"github.com/m00n-arch/remind-stuff-bot/internal/languages"
)

func check(s string) bool {
	// 99.99.9999 99:99
	return regexp.MustCompile(`\d\d.\d\d.\d\d\d\d \d\d:\d\d`).MatchString(s)
}

func (c *Controller) DateHandler(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, languages.MessageEnterText)

	//	if check(update.Message.Text) {
	//		c.alertDB.AddAlert()
	//	}

	err := c.userDB.UpdateState(strconv.FormatInt(update.Message.Chat.ID, 10), db.CreateTextState)
	if err != nil {
		return err
	}

	if _, err := c.bot.Send(msg); err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}
	return nil
}
