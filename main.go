package main

import (
	"github.com/Darcoprogramador/go-notemarkdown/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	// SetupRoutes(app)
	routes.SetupRoutes(app)

	app.Listen(":8000")
}
