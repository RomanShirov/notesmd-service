package db

import (
	"context"
	"github.com/RomanShirov/notesmd-service/internal/models"
	"github.com/teris-io/shortid"
)

func CreateUser(ctx context.Context, user models.AuthUserRequest) (int, error) {
	var uid int
	err = dbConn.QueryRow(context.Background(),
		"INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id",
		user.Username, user.PasswordHash).Scan(&uid)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

func AuthenticateUser(ctx context.Context, username string) (int, string, error) {
	var uid int
	var passwordHash string
	err = dbConn.QueryRow(context.Background(), "SELECT id, password_hash FROM users WHERE username=$1", username).Scan(&uid, &passwordHash)
	if err != nil {
		return 0, "", err
	}

	return uid, passwordHash, nil
}

func GetUsernameFromId(ctx context.Context, uid float64) (string, error) {
	var username string
	err = dbConn.QueryRow(context.Background(), "SELECT username FROM users WHERE id=$1", uid).Scan(&username)
	if err != nil {
		return "", err
	}

	return username, nil
}

func GetNotesBySelectedFolder(ctx context.Context, uid float64, folder string) ([]models.NoteListResponse, error) {
	var notes []models.NoteListResponse
	rows, err := dbConn.Query(context.Background(),
		"SELECT id, uploader_id, folder, title, data FROM notes WHERE uploader_id = $1 AND folder = $2 ORDER BY id", uid, folder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note models.NoteListResponse
		err := rows.Scan(&note.Id, &note.UploaderId, &note.Folder, &note.Title, &note.Data)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func CreateNote(ctx context.Context, uid float64, payload models.CreateNoteRequest) (int, error) {
	var id int
	err = dbConn.QueryRow(context.Background(),
		"INSERT INTO notes (uploader_id, folder, title, data) VALUES ($1, $2, $3, $4) RETURNING id",
		uid, payload.Folder, payload.Title, payload.Data).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateNote(ctx context.Context, uid float64, payload models.UpdateNoteRequest) error {
	_, err = dbConn.Exec(context.Background(),
		"UPDATE notes SET data = $1 WHERE id = $2 AND uploader_id = $3", payload.Data, payload.NoteId, uid)
	if err != nil {
		return err
	}

	return nil
}

func ShareNote(ctx context.Context, uid float64, noteId string) (string, error) {
	shareNoteId, _ := shortid.Generate()
	var id int
	err = dbConn.QueryRow(context.Background(),
		"UPDATE notes SET public_id = $1 WHERE id = $2 AND uploader_id = $3 AND public_id IS NULL RETURNING id",
		shareNoteId, noteId, uid).Scan(&id)
	if err != nil || id == 0 {
		return "", err
	}

	return shareNoteId, nil
}

func DeleteNote(ctx context.Context, uid float64, noteId string) error {
	_, err = dbConn.Exec(context.Background(),
		"DELETE FROM notes WHERE uploader_id = $1 AND id = $2", uid, noteId)
	if err != nil {
		return err
	}

	return nil
}

func GetFolderList(ctx context.Context, uid float64) ([]string, error) {
	folders := []string{}
	rows, err := dbConn.Query(context.Background(),
		"SELECT DISTINCT folder FROM notes WHERE uploader_id = $1", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var folder string
		err := rows.Scan(&folder)
		if err != nil {
			return nil, err
		}
		folders = append(folders, folder)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return folders, nil
}
