package main

import (
	"github.com/RomanShirov/notesmd-service/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		log.Info("Request")
		return c.SendString("Hello, World!")
	})

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartService(app)
	} else {
		utils.GracefulStartService(app)
	}
}
