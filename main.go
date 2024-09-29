package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/Darcoprogramador/go-notemarkdown/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	htmlTemplate "github.com/gofiber/template/html/v2"
)

//go:embed views/*
var viewsfs embed.FS

func main() {

	engine := htmlTemplate.NewFileSystem(http.FS(viewsfs), ".html")
	engine.AddFunc( // add unescape function
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

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
