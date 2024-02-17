package clients

import (
	"log"

	"telegram-bot-api/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(config.Config("API_KEY"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	return bot
}
