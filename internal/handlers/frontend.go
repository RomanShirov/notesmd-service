package handlers

import "github.com/gofiber/fiber/v2"

func SendFrontendStatic(c *fiber.Ctx) error {
	return c.SendFile("./assets/index.html", true)
}
