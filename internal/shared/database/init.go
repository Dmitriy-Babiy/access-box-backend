package database

import (
	"fmt"
	"log"

	"access-box/internal/shared/models"
)

func InitDB() (*DB, error) {
	config := NewConfig()
	
	db, err := NewConnection(config)
	if err != nil {
		return nil, err
	}
	
	log.Println("✅ Соединение с базой данных установлено")
	
	// Инициализируем схему через GORM автомиграции
	if err := InitSchema(db); err != nil {
		return nil, err
	}

	log.Println("✅ Автоматические миграции применены")
	
	return db, nil
}

func RunMigrations(db *DB) error {
	gormInstance := db.DB
	sqlDB, err := gormInstance.DB()
	if err != nil {
		return err
	}

	migrationManager := NewMigrationManager(sqlDB)
	
	if err := migrationManager.RunMigrations("./migrations"); err != nil {
		return err
	}
	
	log.Println("✅ Автоматические миграции применены")
	return nil
}

// InitSchema инициализирует схему базы данных через GORM автомиграции
func InitSchema(db *DB) error {
	// Автоматически создаем таблицы на основе моделей
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("❌ ошибка применения автоматических миграций: %w", err)
	}
	
	return nil
}
