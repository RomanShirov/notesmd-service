package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func InitAuthHandlers(app *fiber.App) {
	auth := app.Group("/api/auth")

	auth.Post("/register", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Route available": true})
	})

	auth.Post("/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Route available": true})
	})
}
