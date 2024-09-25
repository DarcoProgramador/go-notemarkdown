package handlers

import "github.com/gofiber/fiber/v2"

func UploaderHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func RenderMDtoHTMLHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func ListMDFilesHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func CheckMDFileHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
