package repository

import (
	"gorm.io/gorm"
)

type UserRepository interface {
}

type userRepository struct {
	Database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepository{
		Database: database,
	}
}
