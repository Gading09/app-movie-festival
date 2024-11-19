package http

import (
	"github.com/allegro/bigcache/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func ServeHttp(handler Handler, db *gorm.DB, cache *bigcache.BigCache) *fiber.App {

	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())

	RegisterPath(app, handler, cache)

	return app
}
