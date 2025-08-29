package user

import (
	"errors"
)

// UserService интерфейс для бизнес-логики пользователей
type UserService interface {
	CreateUser(email, username, password string) (*User, error)
	GetUserByID(id uint) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(id uint, email, username string) (*User, error)
	DeleteUser(id uint) error
	GetAllUsers() ([]*User, error)
}

// userService реализация сервиса
type userService struct {
	userRepo UserRepository
}

// NewUserService создает новый сервис пользователей
func NewUserService(userRepo UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// CreateUser создает нового пользователя
func (s *userService) CreateUser(email, username, password string) (*User, error) {
	// TODO: Добавить валидацию email и username
	// TODO: Хешировать пароль

	user := NewUser(email, username, password)

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID получает пользователя по ID
func (s *userService) GetUserByID(id uint) (*User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("пользователь не найден")
	}

	return user, nil
}

// GetUserByEmail получает пользователя по email
func (s *userService) GetUserByEmail(email string) (*User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("пользователь не найден")
	}

	return user, nil
}

// UpdateUser обновляет пользователя
func (s *userService) UpdateUser(id uint, email, username string) (*User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("пользователь не найден")
	}

	// TODO: Добавить валидацию
	user.Email = email
	user.Username = username

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser удаляет пользователя
func (s *userService) DeleteUser(id uint) error {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("пользователь не найден")
	}

	return s.userRepo.Delete(id)
}

// GetAllUsers получает всех пользователей
func (s *userService) GetAllUsers() ([]*User, error) {
	return s.userRepo.GetAll()
}
