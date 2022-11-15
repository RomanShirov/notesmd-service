package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func buildConnectionURL() string {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	return fmt.Sprintf("%s:%s", host, port)
}

func StartService(app *fiber.App) {
	connURL := buildConnectionURL()

	if err := app.Listen(connURL); err != nil {
		log.Error("Service is not running! Reason: %v", err)
	}
}

func GracefulStartService(app *fiber.App) {
	connURL := buildConnectionURL()

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Info("Received SIGINT. Shutting down.")
		if err := app.Shutdown(); err != nil {
			log.Error("Service is not shutting down! Reason: %v", err)
		}
	}()

	if err := app.Listen(connURL); err != nil {
		log.Error("Service is not running! Reason: %v", err)
	}
}
