package entity

import "time"

type Comment struct {
	ID          int64         `json:"id" db:"id"`
	CommentText string        `json:"comment_text" db:"comment_text"`
	Author      User          `json:"author" db:"author"`
	CreatedAt   time.Duration `json:"created_at" db:"created_at"`
	IsReview    bool          `json:"is_review" db:"is_review"`
	Grade       int64         `json:"grade" db:"grade"`
	Comment     []*Comment    `json:"comment" db:"comment"`
}
