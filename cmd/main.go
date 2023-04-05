package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/krsnamara/go-docker-postgres-trivia.git/database"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	setupRoutes(app)

	app.Listen(":3000")
}
