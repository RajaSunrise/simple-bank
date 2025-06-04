package main

import (
	"log"

	"github.com/RajaSunrise/simple-bank/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routers.Router(app)

	log.Fatal(app.Listen(":8000"))
}
