package repository

import (
	"tracking_service/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user domain.UserDTO) error
	GetByLogin(username string) (*domain.UserDTO, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user domain.UserDTO) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByLogin(username string) (*domain.UserDTO, error) {
	var user domain.UserDTO
	err := r.db.Where("login = ?", username).First(&user).Error
	return &user, err
}
