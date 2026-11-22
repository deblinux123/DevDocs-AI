package main

import (
	"log"

	"github.com/deblinux123/DevDocs-AI/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":8080"))
}
