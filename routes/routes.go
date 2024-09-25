package routes

import (
	"github.com/Darcoprogramador/go-notemarkdown/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/notes", handlers.UploaderHandler)
	app.Get("/api/notes", handlers.ListMDFilesHandler)
	app.Get("/api/notes/:filename", handlers.RenderMDtoHTMLHandler)
	app.Get("/api/notes/:filename/check", handlers.CheckMDFileHandler)
}
