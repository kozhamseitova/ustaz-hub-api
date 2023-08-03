package entity

import "time"

type Product struct {
	ID          int64         `json:"id" db:"id"`
	User        User          `json:"user" db:"user"`
	Title       string        `json:"title" db:"title"`
	Description string        `json:"description" db:"description"`
	FilePath    string        `json:"file_path" db:"file_path"`
	UploadedAt  time.Duration `json:"uploaded_at" db:"uploaded_at"`
}
