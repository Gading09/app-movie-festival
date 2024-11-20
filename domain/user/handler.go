package user

import (
	"movie-festival/delivery/http/middleware"
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
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
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

func (handler userHandler) Login(c *fiber.Ctx) error {
	login := new(model.ReqLogin)
	if err := c.BodyParser(login); err != nil {
		err = e.New(constant.StatusBadRequest, constant.ErrInvalidRequest, err)
		return response.ResponseError(c, err)
	}
	if err, check := validator.Validation(login); check {
		err = e.New(constant.StatusBadRequest, constant.ErrValidator, err)
		return response.ResponseError(c, err)
	}
	token, err := handler.Feature.LoginFeature(login)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseOK(c, http.StatusCreated, constant.GetSuccess, token)
}

func (handler userHandler) Logout(c *fiber.Ctx) error {
	UserContext := c.UserContext()
	payloadToken := UserContext.Value(constant.DATA_TOKEN).(middleware.DataUserToken)
	userId := payloadToken.Profile.Id
	if userId == "" {
		err := e.New(constant.StatusBadRequest, constant.ErrAuth, nil)
		return response.ResponseError(c, err)
	}
	err := handler.Feature.LogoutFeature(userId)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseOK(c, http.StatusOK, constant.Success, nil)
}
