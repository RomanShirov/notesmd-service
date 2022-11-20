package models

type AuthUserRequest struct {
	Username     string
	PasswordHash string
}

type NoteListResponse struct {
	Id         int         `json:"id"`
	UploaderId int         `json:"uploader_id"`
	Folder     string      `json:"folder"`
	Title      string      `json:"title"`
	Data       interface{} `json:"data"`
}

type CreateNoteRequest struct {
	Folder string      `json:"folder" form:"folder"`
	Title  string      `json:"title" form:"title"`
	Data   interface{} `json:"data" form:"data"`
}

type UpdateNoteRequest struct {
	NoteId int         `json:"note_id" form:"note_id"`
	Data   interface{} `json:"data" form:"data"`
}
