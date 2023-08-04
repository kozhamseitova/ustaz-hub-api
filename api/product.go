package api

import "time"

type CreateProductRequest struct {
	ID          int64     `json:"id" db:"id"`
	UserId      int64     `json:"user_id" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	FilePath    string    `json:"file_path" db:"file_path"`
	UploadedAt  time.Time `json:"uploaded_at" db:"uploaded_at"`
}
