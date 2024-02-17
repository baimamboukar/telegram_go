package repositories

import (
	"time"

	"telegram-bot-api/src/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetTask(update tgbotapi.Update) error {
	member := models.Member{
		FullName: update.Message.AuthorSignature,
		SocialMedia: map[string]string{
			"whatsapp": "+237690535759",
			"facebook": "facebook.com/baimamboukar",
			"x":        "https://x.com/baimamjj",
			"linkedin": "https://linkedin.com/in/baimamboukar",
		},
		Title:      "Events Manager",
		Age:        22,
		Bio:        "Software Engineer and Cloud Native Apps dev",
		DateJoined: time.Time{},
	}

	if result := DB.Create(&member); result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteTask(taskId string) error {
	if result := DB.Where("id = ?", taskId).Delete(&models.Member{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllTasks(chatId int64) ([]models.Member, error) {
	var tasks []models.Member
	if result := DB.Where("chat_id = ?", chatId).Find(&tasks); result.Error != nil {
		return tasks, result.Error
	}
	return tasks, nil
}
