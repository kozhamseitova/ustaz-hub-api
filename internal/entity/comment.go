package entity

import (
	"time"
)

type Comment struct {
	ID          int64  `json:"id" db:"id"`
	ParentID    int64  `json:"parent_id" db:"parent_id"`
	CommentText string `json:"comment_text" db:"comment_text"`
	//Author      User       `json:"author" db:"author"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	IsReview  bool       `json:"is_review" db:"is_review"`
	Grade     string     `json:"grade" db:"grade"`
	Comments  []*Comment `json:"comments" db:"comments"`
}

type CreateComment struct {
	ID          int64  `json:"id" db:"id"`
	ParentID    int64  `json:"paren_id" db:"parent_id"`
	CommentText string `json:"comment_text" db:"comment_text"`
	AuthorID    int64  `json:"author_id" db:"author_id"`
	IsReview    bool   `json:"is_review" db:"is_review"`
	ProductID   int64  `json:"product_id" db:"product_id"`
	Grade       int64  `json:"grade" db:"grade"`
	PostID      int64  `json:"post_id" db:"post_id"`
}
