package handlers

import (
	"context"
	"fmt"
	"github.com/RomanShirov/notesmd-service/internal/crypto"
	db "github.com/RomanShirov/notesmd-service/internal/database"
	"github.com/RomanShirov/notesmd-service/internal/models"
	"github.com/gofiber/fiber/v2"
	"os"
)

func InitNotesAPI(app *fiber.App) {
	notes := app.Group("/api/notes")

	notes.Get("/:folder", func(c *fiber.Ctx) error {
		requestFolder := c.Params("folder")
		userId := crypto.GetUserIdFromToken(c)
		notes, err := db.GetNotesBySelectedFolder(context.Background(), userId, requestFolder)
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.JSON(notes)
	})

	notes.Put("/", func(c *fiber.Ctx) error {
		encodedRequest := new(models.CreateNoteRequest)
		userId := crypto.GetUserIdFromToken(c)

		if err := c.BodyParser(encodedRequest); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		noteId, err := db.CreateNote(context.Background(), userId, *encodedRequest)

		if err != nil && noteId == 0 {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{"id": noteId})
	})

	notes.Put("/share/:note_id", func(c *fiber.Ctx) error {
		requestNoteId := c.Params("note_id")
		userId := crypto.GetUserIdFromToken(c)
		username, err := db.GetUsernameFromId(context.Background(), userId)

		sharedNoteKey, err := db.ShareNote(context.Background(), userId, requestNoteId)
		if err != nil && sharedNoteKey == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		sharedNoteURL := fmt.Sprintf("%s/shared/%s/%s", os.Getenv("PUBLIC_SERVER_URL"), username, sharedNoteKey)

		return c.JSON(fiber.Map{"public_url": sharedNoteURL})
	})

	notes.Patch("/", func(c *fiber.Ctx) error {
		encodedRequest := new(models.UpdateNoteRequest)
		userId := crypto.GetUserIdFromToken(c)

		if err := c.BodyParser(encodedRequest); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err := db.UpdateNote(context.Background(), userId, *encodedRequest)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{"success": true})
	})

	notes.Delete("/:note_id", func(c *fiber.Ctx) error {
		requestNoteId := c.Params("note_id")
		userId := crypto.GetUserIdFromToken(c)

		err := db.DeleteNote(context.Background(), userId, requestNoteId)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{"success": true})
	})

}
