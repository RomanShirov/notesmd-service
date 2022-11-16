package db

import (
	"context"
	"github.com/RomanShirov/notesmd-service/internal/models"
)

func CreateUser(ctx context.Context, user models.AuthUserRequest) error {
	_, err := dbConn.Exec(context.Background(),
		"INSERT INTO users (email, password_hash) VALUES ($1, $2)",
		user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func AuthenticateUser(ctx context.Context, email string) (string, error) {
	var passwordHash string
	err = dbConn.QueryRow(context.Background(), "SELECT password_hash FROM users WHERE email=$1", email).Scan(&passwordHash)
	if err != nil {
		return "", err
	}

	return passwordHash, nil
}
