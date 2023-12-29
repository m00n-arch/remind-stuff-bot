package controller

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m00n-arch/remind-stuff-bot/internal/buttons"
	"github.com/m00n-arch/remind-stuff-bot/internal/db"
)

func (c *Controller) ExistingRemHandler(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Все активные напоминания представлены ниже:")

	err := c.userDB.UpdateState(strconv.FormatInt(update.Message.Chat.ID, 10), db.StartState)
	if err != nil {
		return err
	}

	res, err := c.alertDB.GetAlerts(strconv.Itoa(int(update.Message.From.ID)))
	if err != nil {
		return err
	}

	str := ""
	for i := range res {
		str += fmt.Sprintf("%s - %s\n", res[i].Content, res[i].Date.Format("02.01.2006 15:04"))
	}
	msg.Text = str

	msg.ReplyMarkup = buttons.Cancel

	if _, err := c.bot.Send(msg); err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}
	return nil
}
