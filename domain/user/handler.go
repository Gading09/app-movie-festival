package user

import (
	"movie-festival/domain/user/feature"
	"movie-festival/domain/user/model"
	"movie-festival/helper/constant"
	"movie-festival/helper/response"
	e "movie-festival/helper/response/error"
	validator "movie-festival/helper/validator"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	RegisterUser(c *fiber.Ctx) error
}

type userHandler struct {
	Feature feature.UserFeature
}

func NewUserHandler(feature feature.UserFeature) UserHandler {
	return &userHandler{
		Feature: feature,
	}
}

func (handler userHandler) RegisterUser(c *fiber.Ctx) error {
	user := new(model.ReqUser)
	if err := c.BodyParser(user); err != nil {
		err = e.New(constant.StatusBadRequest, constant.ErrInvalidRequest, err)
		return response.ResponseError(c, err)
	}
	if err, check := validator.Validation(user); check {
		err = e.New(constant.StatusBadRequest, constant.ErrValidator, err)
		return response.ResponseError(c, err)
	}
	err := handler.Feature.RegisterUserFeature(user)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseOK(c, http.StatusCreated, constant.CreateSuccess, nil)
}
