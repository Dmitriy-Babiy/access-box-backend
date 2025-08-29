package database

import (
	"fmt"
	"log"

	"access-box/internal/shared/models"
)

// InitDB инициализирует базу данных и возвращает подключение
func InitDB() (*DB, error) {
	config := NewConfig()
	
	// Создаем базу данных если она не существует
	if err := CreateDatabase(config); err != nil {
		return nil, err
	}
	
	// Подключаемся к созданной базе данных
	db, err := NewConnection(config)
	if err != nil {
		return nil, err
	}
	
	// Инициализируем схему через GORM автомиграции
	if err := InitSchema(db); err != nil {
		return nil, err
	}
	
	log.Println("✅ База данных успешно инициализирована")
	return db, nil
}

// CreateDatabase создает базу данных если она не существует
func CreateDatabase(config *Config) error {
	// Создаем временное подключение к postgres для создания БД
	tempConfig := &Config{
		DBHost:     config.DBHost,
		DBPort:     config.DBPort,
		DBUser:     config.DBUser,
		DBPassword: config.DBPassword,
		DBName:     "postgres", // Подключаемся к системной БД
	}

	tempDB, err := NewConnection(tempConfig)
	if err != nil {
		return err
	}
	defer tempDB.Close()

	// Проверяем существование БД
	var count int64
	err = tempDB.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", config.DBName).Scan(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		log.Printf("📝 Создаем базу данных '%s'...", config.DBName)
		err = tempDB.Exec("CREATE DATABASE " + config.DBName).Error
		if err != nil {
			return err
		}
		log.Printf("✅ База данных '%s' успешно создана", config.DBName)
	} else {
		log.Printf("ℹ️ База данных '%s' уже существует", config.DBName)
	}

	return nil
}

// RunMigrations запускает миграции базы данных
func RunMigrations(db *DB) error {
	// Получаем SQL подключение для Goose
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}

	// Создаем менеджер миграций
	migrationManager := NewMigrationManager(sqlDB)
	
	// Запускаем миграции
	if err := migrationManager.RunMigrations("./migrations"); err != nil {
		return err
	}
	
	log.Println("✅ Миграции успешно применены")
	return nil
}

// InitSchema инициализирует схему базы данных через GORM автомиграции
func InitSchema(db *DB) error {
	log.Println("🔧 Инициализируем схему базы данных...")
	
	// Автоматически создаем таблицы на основе моделей
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("failed to auto-migrate User model: %w", err)
	}
	
	log.Println("✅ Схема базы данных инициализирована")
	return nil
}
