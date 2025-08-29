package user

import (
	"access-box/internal/shared/models"
	"time"
)

// User представляет сущность пользователя
type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser создает нового пользователя
func NewUser(email, username, password string) *User {
	return &User{
		Email:     email,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// ToModel конвертирует entity в GORM модель
func (u *User) ToModel() *models.User {
	return &models.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// FromModel конвертирует GORM модель в entity
func FromModel(model *models.User) *User {
	return &User{
		ID:        model.ID,
		Email:     model.Email,
		Username:  model.Username,
		Password:  model.Password,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
