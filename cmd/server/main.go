package main

import (
	"log"

	"github.com/deblinux123/DevDocs-AI/internal/config"
	"github.com/deblinux123/DevDocs-AI/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
