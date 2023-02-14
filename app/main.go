package main

import (
	"github.com/diiegoburiti/go-book/app/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	configs.ConnectDB()

	app.Listen(":6000")
}
