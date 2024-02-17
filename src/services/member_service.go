package services

import (
	"telegram-bot-api/src/keyboards"
	"telegram-bot-api/src/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Hi, here you can create todos for your todolist."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SetTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Please, write todo."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func CreateMemberCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Member successfully created"

	err := repositories.CreateMember(update)
	if err != nil {
		text = "Couldnt set task"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func ShowAllTasks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

}
