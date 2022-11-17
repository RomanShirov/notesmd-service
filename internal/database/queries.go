package db

import (
	"context"
	"github.com/RomanShirov/notesmd-service/internal/models"
)

func CreateUser(ctx context.Context, user models.AuthUserRequest) (int, error) {
	var uid int
	err = dbConn.QueryRow(context.Background(),
		"INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id",
		user.Email, user.PasswordHash).Scan(&uid)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

func AuthenticateUser(ctx context.Context, email string) (int, string, error) {
	var uid int
	var passwordHash string
	err = dbConn.QueryRow(context.Background(), "SELECT id, password_hash FROM users WHERE email=$1", email).Scan(&uid, &passwordHash)
	if err != nil {
		return 0, "", err
	}

	return uid, passwordHash, nil
}
