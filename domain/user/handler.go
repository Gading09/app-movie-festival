package user

import (
	"movie-festival/domain/user/feature"
)

type UserHandler interface {
}

type userHandler struct {
	Feature feature.UserFeature
}

func NewUserHandler(feature feature.UserFeature) UserHandler {
	return &userHandler{
		Feature: feature,
	}
}
