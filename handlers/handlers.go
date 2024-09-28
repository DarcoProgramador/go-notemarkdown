package handlers

import (
	"log/slog"
	"strings"

	"github.com/Darcoprogramador/go-notemarkdown/storage"
	"github.com/gofiber/fiber/v2"
)

func UploaderHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if strings.ToLower(strings.Split(file.Filename, ".")[1]) != "md" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "File must be a markdown file",
		})
	}

	slog.Info("File recived", slog.String("file", file.Filename))
	path, err := storage.SaveFile(file, "./storage/temp")

	if err != nil {
		slog.Error("Error saving file", slog.String("file", file.Filename), slog.String("error", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "File uploaded successfully",
		"path":    path,
	})
}

func RenderMDtoHTMLHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func ListMDFilesHandler(c *fiber.Ctx) error {
	fileNames, err := storage.GetFiles("./storage/temp")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"files": fileNames,
	})
}

func CheckMDFileHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
