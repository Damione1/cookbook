package main

import (
	database "github.com/Damione1/portfolio-playground/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
