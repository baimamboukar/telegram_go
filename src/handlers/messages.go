package handlers

import (
	"telegram-bot-api/src/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	services.CreateMemberCallback(bot, update)
}
