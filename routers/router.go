package routers

import (
	"github.com/RajaSunrise/simple-bank/db/sqlc"
	"github.com/RajaSunrise/simple-bank/handlers"
	"github.com/gofiber/fiber/v2"
)


func Router(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	queries := sqlc.New()
	UserHandler := handlers.NewUsersHandler()
	users := v1.Group("/users")
	users.Post("/", UserHandler.CreateUsers)
}
