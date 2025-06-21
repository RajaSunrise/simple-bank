package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupMiddleware(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
