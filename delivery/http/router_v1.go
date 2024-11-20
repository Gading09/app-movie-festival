package http

import (
	"github.com/allegro/bigcache/v3"
	"github.com/gofiber/fiber/v2"
)

func RegisterPath(app *fiber.App, h Handler, cache *bigcache.BigCache) {
	user := app.Group("/user")
	{
		user.Post("/registration", h.UserHandler.RegisterUser)
		user.Post("/login", h.UserHandler.Login)
	}
}
