package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func NewConnection(config *Config) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Логируем SQL запросы
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия соединения с БД: %w", err)
	}

	// Получаем *sql.DB для настройки пула соединений
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("ошибка получения *sql.DB: %w", err)
	}

	// Настройка пула соединений
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)

	log.Println("✅ Успешно подключились к PostgreSQL через GORM")
	return &DB{db}, nil
}

func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *DB) Ping() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func (db *DB) IsConnected() bool {
	return db.Ping() == nil
}
