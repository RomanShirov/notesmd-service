package handlers

import (
	"context"
	"github.com/RomanShirov/notesmd-service/internal/crypto"
	db "github.com/RomanShirov/notesmd-service/internal/database"
	"github.com/RomanShirov/notesmd-service/internal/models"
	"github.com/gofiber/fiber/v2"
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
