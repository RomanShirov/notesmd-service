package main

import (
	"fmt"
	"github.com/RomanShirov/notesmd-service/internal/database"
	"github.com/RomanShirov/notesmd-service/internal/handlers"
	"github.com/RomanShirov/notesmd-service/internal/utils"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbConnURL := fmt.Sprintf("postgres://%s:%s@%s/notesdb?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_URL"))

	if os.Getenv("MIGRATIONS") == "true" {
		log.Info("Starting migrations")
		db.RollupMigrations(dbConnURL)
		if err != nil {
			log.Fatalf("Database migration error: %v", err)
		}
	}

	db.InitDatabase(dbConnURL)

	app := fiber.New()

	app.Use(cors.New())

	prometheus := fiberprometheus.New("notes-app-metrics")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Static("/", "./assets", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: 3600 * time.Second,
		MaxAge:        3600,
	})

	handlers.InitAuthHandlers(app)

	app.Get("/", handlers.SendFrontendStatic)

	// Force frontend application for prevent 404 fallthrough
	app.Get("/auth", handlers.SendFrontendStatic)
	app.Get("/shared/*/*", handlers.SendFrontendStatic)
	app.Get("/api/notes/shared/:shared_id", handlers.GetSharedNote)

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
