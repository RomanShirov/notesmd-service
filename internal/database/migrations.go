package db

import (
	"database/sql"
	"embed"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func RollupMigrations(dbConnURL string) {

	goose.SetBaseFS(embedMigrations)

	time.Sleep(5 * time.Second)
	db, err := sql.Open("postgres", dbConnURL)
	if err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	if os.Getenv("CLEAN_DEPLOY") == "true" {
		log.Info("CLEAN_DEPLOY: Rolling back old migrations")
		if err := goose.Down(db, "migrations"); err != nil {
			log.Fatalf("Migration error: %v", err)
		}
	}
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
}
