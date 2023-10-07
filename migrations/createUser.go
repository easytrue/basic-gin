package migrations

import (
	"basicGin/model"
	"gorm.io/gorm"
)

func CreateUserTable(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
