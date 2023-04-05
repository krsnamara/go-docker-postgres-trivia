package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krsnamara/go-docker-postgres-trivia.git/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hollow, World/")
	})

	app.Listen(":3000")
}
