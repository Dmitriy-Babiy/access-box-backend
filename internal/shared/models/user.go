package models

import (
	"time"

	"gorm.io/gorm"
)

// User представляет модель пользователя для GORM
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null"`
	Password  string         `json:"-" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName указывает имя таблицы для GORM
func (User) TableName() string {
	return "users"
}
