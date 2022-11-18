package handlers

import (
	"context"
	"github.com/RomanShirov/notesmd-service/internal/crypto"
	db "github.com/RomanShirov/notesmd-service/internal/database"
	"github.com/gofiber/fiber/v2"
)

func InitFoldersAPI(app *fiber.App) {
	folders := app.Group("/api/folders")

	folders.Get("/", func(c *fiber.Ctx) error {
		userId := crypto.GetUserIdFromToken(c)
		folderList, err := db.GetFolderList(context.Background(), userId)
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.JSON(folderList)
	})
}
