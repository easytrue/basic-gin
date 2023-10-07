package migrations

import (
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {

	// 用户表
	CreateUserTable(db)
}
