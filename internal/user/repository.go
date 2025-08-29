package user

import (
	"access-box/internal/shared/database"
	"access-box/internal/shared/models"
)

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id uint) error
	GetAll() ([]*User, error)
}

// userRepository реализация репозитория
type userRepository struct {
	db *database.DB
}

// NewUserRepository создает новый репозиторий пользователей
func NewUserRepository(db *database.DB) UserRepository {
	return &userRepository{db: db}
}

// Create создает нового пользователя
func (r *userRepository) Create(user *User) error {
	model := user.ToModel()
	err := r.db.Create(model).Error
	if err != nil {
		return err
	}
	user.ID = model.ID
	user.CreatedAt = model.CreatedAt
	user.UpdatedAt = model.UpdatedAt
	return nil
}

// GetByID получает пользователя по ID
func (r *userRepository) GetByID(id uint) (*User, error) {
	var model models.User
	err := r.db.First(&model, id).Error
	if err != nil {
		return nil, err
	}
	return FromModel(&model), nil
}

// GetByEmail получает пользователя по email
func (r *userRepository) GetByEmail(email string) (*User, error) {
	var model models.User
	err := r.db.Where("email = ?", email).First(&model).Error
	if err != nil {
		return nil, err
	}
	return FromModel(&model), nil
}

// Update обновляет пользователя
func (r *userRepository) Update(user *User) error {
	model := user.ToModel()
	return r.db.Save(model).Error
}

// Delete удаляет пользователя (soft delete)
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// GetAll получает всех пользователей
func (r *userRepository) GetAll() ([]*User, error) {
	var modelUsers []*models.User
	err := r.db.Find(&modelUsers).Error
	if err != nil {
		return nil, err
	}
	
	users := make([]*User, len(modelUsers))
	for i, model := range modelUsers {
		users[i] = FromModel(model)
	}
	return users, nil
}
