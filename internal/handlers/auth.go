package handlers

import (
	"context"
	"github.com/RomanShirov/notesmd-service/internal/crypto"
	db "github.com/RomanShirov/notesmd-service/internal/database"
	"github.com/RomanShirov/notesmd-service/internal/models"
	"github.com/gofiber/fiber/v2"
)

func InitAuthHandlers(app *fiber.App) {
	auth := app.Group("/api/auth")

	auth.Post("/register", func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		user := models.AuthUserRequest{
			Email:        email,
			PasswordHash: crypto.GeneratePasswordHash(password),
		}
		err := db.CreateUser(context.Background(), user)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		token := crypto.GenerateUserToken(email)
		return c.JSON(fiber.Map{"access_token": token})
	})

	auth.Post("/login", func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		passwordHash, err := db.AuthenticateUser(context.Background(), email)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if crypto.ComparePasswords(password, passwordHash) {
			token := crypto.GenerateUserToken(email)
			return c.JSON(fiber.Map{"access_token": token})
		}

		return nil
	})
}
