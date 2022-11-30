package handlers

import (
	"context"
	"github.com/RomanShirov/notesmd-service/internal/crypto"
	db "github.com/RomanShirov/notesmd-service/internal/database"
	"github.com/RomanShirov/notesmd-service/internal/models"
	"github.com/gofiber/fiber/v2"
)

func InitAuthHandlers(app *fiber.App) {
	app.Post("/register", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		user := models.AuthUserRequest{
			Username:     username,
			PasswordHash: crypto.GeneratePasswordHash(password),
		}
		uid, err := db.CreateUser(context.Background(), user)
		if err != nil && uid == 0 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		token := crypto.GenerateUserToken(uid)
		return c.JSON(fiber.Map{
			"access_token": token,
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		uid, passwordHash, err := db.AuthenticateUser(context.Background(), username)
		if err != nil && uid == 0 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if crypto.ComparePasswords(password, passwordHash) {
			token := crypto.GenerateUserToken(uid)
			return c.JSON(fiber.Map{"access_token": token})
		}

		return nil
	})
}
