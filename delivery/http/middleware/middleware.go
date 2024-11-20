package middleware

import (
	"strings"

	"github.com/allegro/bigcache/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"movie-festival/helper/constant"
	"movie-festival/helper/context"
	"movie-festival/helper/response"
	e "movie-festival/helper/response/error"
)

type Profile struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"isAdmin"`
}

type DataUserToken struct {
	Profile Profile `json:"profile"`
	Exp     int     `json:"exp"`
	Sub     string  `json:"sub"`
	jwt.RegisteredClaims
}

func CheckTokenExpire(cache *bigcache.BigCache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.CreateContext()
		ctx = c.UserContext()
		var dataUserToken DataUserToken
		token := c.Get("Authorization")
		if token == "" {
			err := e.New(constant.StatusBadRequest, constant.ErrAuth, nil)
			return response.ResponseError(c, err)
		}
		authToken := strings.Split(token, " ")[1]
		if _, err := jwt.ParseWithClaims(authToken, &dataUserToken, func(token *jwt.Token) (interface{}, error) {
			return []byte{}, nil
		}); err != nil {
			err := e.New(constant.StatusBadRequest, constant.ErrAuth, nil)
			return response.ResponseError(c, err)
		}
		entry, err := cache.Get(dataUserToken.Profile.Id)
		if err != nil || string(entry) != authToken {
			if strings.Contains(err.Error(), "not found") {
				err := e.New(constant.StatusBadRequest, constant.ErrAuth, nil)
				return response.ResponseError(c, err)
			}
		}
		ctx = context.SetTokenStructToContext(ctx, constant.DATA_TOKEN, dataUserToken)
		c.SetUserContext(ctx)
		c.Next()
		return nil
	}
}

func IsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.CreateContext()

		UserContext := c.UserContext()
		payloadToken := UserContext.Value(constant.DATA_TOKEN).(DataUserToken)
		isAdmin := payloadToken.Profile.IsAdmin
		if !isAdmin {
			err := e.New(constant.StatusBadRequest, constant.ErrAuth, nil)
			return response.ResponseError(c, err)
		}
		c.SetUserContext(ctx)
		c.Next()
		return nil
	}
}
