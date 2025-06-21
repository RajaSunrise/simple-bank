package main

import (
	"fmt"
	"log"

	"github.com/RajaSunrise/simple-bank/routers"
	"github.com/RajaSunrise/simple-bank/utils"
	"github.com/gofiber/fiber/v2"
)

// @title Fiber Swagger
// @version 1.0
// @description website example bank with fiber
// @host localhost:8000
// @BasePath /
func main() {
	app := fiber.New(fiber.Config{
		AppName: utils.AppName,
		Prefork: utils.AppStatus,
	})

	routers.SetupMiddleware(app)
	routers.SetupPublicRouters(app)
	routers.SetupPrivateRouters(app)

	host := fmt.Sprintf("%s:%d", utils.AppHost, utils.AppPort)
	log.Fatal(app.Listen(host))
}
