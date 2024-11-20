package feature

import (
	"movie-festival/domain/user/model"
	"movie-festival/domain/user/repository"
	"movie-festival/helper/constant"
	j "movie-festival/helper/jwt"
	e "movie-festival/helper/response/error"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/google/uuid"
)

type UserFeature interface {
	RegisterUserFeature(request *model.ReqUser) (err error)
	LoginFeature(request *model.ReqLogin) (token model.ResLogin, err error)
}

type userFeature struct {
	Repository repository.UserRepository
	Cache      *bigcache.BigCache
}

func NewUserFeature(repository repository.UserRepository, cache *bigcache.BigCache) UserFeature {
	return &userFeature{
		Repository: repository,
		Cache:      cache,
	}
}

func (feature userFeature) RegisterUserFeature(request *model.ReqUser) (err error) {
	hashPassword, err := j.HashPassword(request.Password)
	if err != nil {
		err = e.New(constant.StatusBadRequest, constant.ErrDefault, err)
		return
	}
	payload := model.User{
		Id:        uuid.New().String(),
		Username:  request.Username,
		Password:  hashPassword,
		Email:     request.Email,
		CreatedAt: time.Now(),
		IsAdmin:   *request.IsAdmin,
	}
	return feature.Repository.RegisterUserRepository(payload)
}

func (feature userFeature) LoginFeature(request *model.ReqLogin) (res model.ResLogin, err error) {
	user, err := feature.Repository.GetUserByEmailRepository(request.Email)
	if err != nil {
		return
	}
	err = j.ComparePassword(user.Password, request.Password)
	if err != nil {
		err = e.New(constant.StatusBadRequest, constant.ErrAuth, err)
		return
	}
	token, err := j.GenerateToken(map[string]interface{}{
		"profile": map[string]interface{}{
			"id":       user.Id,
			"username": user.Username,
			"email":    user.Email,
			"isAdmin":  user.IsAdmin,
		},
	}, time.Duration(1)*time.Hour)
	if err != nil {
		err = e.New(constant.StatusBadRequest, constant.ErrDefault, err)
		return
	}
	err = feature.Cache.Set(user.Id, []byte(token))
	if err != nil {
		err = e.New(constant.StatusBadRequest, constant.ErrDefault, err)
		return
	}
	return model.ResLogin{
		Token: token,
	}, nil
}
