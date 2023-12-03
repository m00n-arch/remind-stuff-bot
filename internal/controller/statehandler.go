package controller

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Controller) StateHandler(update tgbotapi.Update) error {
	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите интересующее вас действие")

	err := c.userDB.UpdateState(strconv.FormatInt(update.Message.Chat.ID, 10), "createState")
	if err != nil {
		return err
	}

	return nil
	//	msg.Text = "Введите дату и время для напоминания"
}
