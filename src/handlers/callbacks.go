package handlers

import (
	"telegram-bot-api/src/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmd, taskId := "lol", "lol"
	//utils.GetKeyValue(update.CallbackQuery.Data)
	switch {
	case cmd == "detele_member":
		services.DeleteMemberCallBack(bot, update, taskId)
	}
}
