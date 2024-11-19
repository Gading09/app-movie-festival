package feature

import (
	"movie-festival/domain/user/repository"
)

type UserFeature interface {
}

type userFeature struct {
	Repository repository.UserRepository
}

func NewUserFeature(repository repository.UserRepository) UserFeature {
	return &userFeature{
		Repository: repository,
	}
}
