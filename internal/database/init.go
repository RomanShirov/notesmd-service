package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"os"
)

var dbConn *pgxpool.Pool

var err error

func InitDatabase(dbConnURL string) {
	log.Info("Starting connection")
	dbConn, err = pgxpool.New(context.Background(), dbConnURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	log.Info("Connection successful")
}
