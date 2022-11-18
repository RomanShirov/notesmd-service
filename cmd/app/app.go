package main

import (
	"github.com/RomanShirov/notesmd-service/internal/database"
	"github.com/RomanShirov/notesmd-service/internal/handlers"
	"github.com/RomanShirov/notesmd-service/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if os.Getenv("MIGRATIONS") == "true" {
		log.Info("Starting migrations")
		db.RollupMigrations()
		if err != nil {
			log.Fatalf("Database migration error: %v", err)
		}
	}

	db.InitDatabase()

	app := fiber.New()

	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	handlers.InitAuthHandlers(app)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}))

	handlers.InitNotesAPI(app)
	handlers.InitFoldersAPI(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartService(app)
	} else {
		utils.GracefulStartService(app)
	}
}
