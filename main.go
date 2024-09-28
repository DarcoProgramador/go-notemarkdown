package main

import (
	"github.com/Darcoprogramador/go-notemarkdown/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	htmlTemplate "github.com/gofiber/template/html/v2"
)

func main() {

	engine := htmlTemplate.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	// SetupRoutes(app)
	routes.SetupRoutes(app)

	app.Listen(":8000")
}
