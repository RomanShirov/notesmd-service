package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func InitNotesAPI(app *fiber.App) {
	auth := app.Group("/api/notes")

	auth.Get("/:folder", func(c *fiber.Ctx) error {
		requestFolder := c.Params("folder")
		return c.JSON(fiber.Map{"folder": requestFolder})
	})

	auth.Put("/", func(c *fiber.Ctx) error {
		requestFolder := c.Params("folder")
		return c.JSON(fiber.Map{"folder": requestFolder})
	})

}
