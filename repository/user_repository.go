package repository

import (
	"go-rest-api-todo/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (repository *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := repository.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (repository *userRepository) CreateUser(user *model.User) error {
	if err := repository.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}