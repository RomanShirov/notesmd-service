package db

import (
	"database/sql"
	"embed"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
	"os"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func RollupMigrations() {

	goose.SetBaseFS(embedMigrations)

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
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

	if os.Getenv("CLEAR_DEPLOY") == "true" {
		log.Info("CLEAR_DEPLOY: Rolling back migrations")
		if err := goose.Down(db, "migrations"); err != nil {
			log.Fatalf("Migration error: %v", err)
		}
	}
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
}