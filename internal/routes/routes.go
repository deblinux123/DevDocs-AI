package routes

import (
	"github.com/deblinux123/DevDocs-AI/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/health", handlers.Health)
}
