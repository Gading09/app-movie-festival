package http

import (
	"movie-festival/delivery/http/middleware"

	"github.com/allegro/bigcache/v3"
	"github.com/gofiber/fiber/v2"
)

func RegisterPath(app *fiber.App, h Handler, cache *bigcache.BigCache) {
	user := app.Group("/user")
	{
		user.Post("/registration", h.UserHandler.RegisterUser)
		user.Post("/login", h.UserHandler.Login)
		user.Post("/logout", middleware.CheckTokenExpire(cache), h.UserHandler.Logout)
	}

	admin := app.Group("/admin")
	{
		admin.Post("/movie", middleware.CheckTokenExpire(cache), middleware.IsAdmin(), h.MovieHandler.CreateMovie)
		admin.Put("/movie/:movieId", middleware.CheckTokenExpire(cache), middleware.IsAdmin(), h.MovieHandler.UpdateMovie)
	}
}
