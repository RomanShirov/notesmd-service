package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
	"os"
)

var dbConn *pgx.Conn

var err error

func InitDatabase(dbConnURL string) {
	log.Info("Starting connection")
	dbConn, err = pgx.Connect(context.Background(), dbConnURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	log.Info("Connection successful")
}
