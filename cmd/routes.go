package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krsnamara/go-docker-postgres-trivia.git/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Get("/fact", handlers.NewFact)

	app.Post("/fact", handlers.CreateFact)
}
