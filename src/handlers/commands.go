package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot-api/src/services"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		services.Start(bot, update)
	case "set_todo":
		services.SetTask(bot, update)
	case "delete_todo":
		services.DeleteTask(bot, update)
	case "show_all_todos":
		services.ShowAllTasks(bot, update)
	}
}
