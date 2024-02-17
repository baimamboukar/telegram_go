package repositories

import (
	"gorm.io/gorm"

	"telegram-bot-api/src/database"
)

var DB *gorm.DB = database.Init()
