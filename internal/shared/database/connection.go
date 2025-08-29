package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewConnection(config *Config) (*DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("❌ ошибка открытия соединения с БД: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("❌ ошибка получения *sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)

	return &DB{db}, nil
}

