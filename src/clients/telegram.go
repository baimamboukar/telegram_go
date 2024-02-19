package clients

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot-api/config"
)

func Init() *tgbotapi.BotAPI {

	bot, err := tgbotapi.NewBotAPI(config.Config("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	return bot
}
