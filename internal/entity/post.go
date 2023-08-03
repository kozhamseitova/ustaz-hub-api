package entity

import "time"

type Post struct {
	ID        string        `json:"id" db:"id"`
	User      User          `json:"user" db:"user"`
	PostText  string        `json:"post_text" db:"post_text"`
	CreatedAt time.Duration `json:"created_at" db:"created_at"`
}
