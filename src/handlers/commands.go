package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot-api/src/services"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		services.Start(bot, update)
	case "create_member":
		services.CreateMemberCallback(bot, update)
	case "delete_member":
		services.DeleteMember(bot, update)
	case "get_all_members":
		services.GetAllMembers(bot, update)
	}
}
