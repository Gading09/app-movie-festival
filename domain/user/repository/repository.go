package repository

import (
	"movie-festival/domain/user/model"
	"movie-festival/helper/constant"
	e "movie-festival/helper/response/error"
	"strings"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUserRepository(payload model.User) (err error)
	GetUserByEmailRepository(email string) (user model.User, err error)
}

type userRepository struct {
	Database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepository{
		Database: database,
	}
}

func (repo userRepository) RegisterUserRepository(payload model.User) (err error) {
	if err = repo.Database.Create(payload).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			err = e.New(constant.StatusInternalServerError, constant.ErrAlreadyExist, err)
			return
		}
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return err
	}
	return nil
}

func (repo userRepository) GetUserByEmailRepository(email string) (user model.User, err error) {
	if err = repo.Database.Where("email = ?", email).Find(&user).Error; err != nil {
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return
	}
	return
}
