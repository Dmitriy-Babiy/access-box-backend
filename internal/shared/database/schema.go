package database

import (
	"access-box/internal/shared/models"
	"log"
)

// InitSchema создает таблицы в базе данных используя GORM AutoMigrate
func (db *DB) InitSchema() error {
	log.Println("🔧 Инициализация схемы базы данных через GORM...")

	// Автоматическая миграция таблиц
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	log.Println("✅ Схема базы данных успешно инициализирована через GORM")
	return nil
}
