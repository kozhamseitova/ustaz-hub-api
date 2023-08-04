package api

type CreatePostRequest struct {
	ID       int64  `json:"id" db:"id"`
	UserId   int64  `json:"user_id" db:"user_id"`
	PostText string `json:"post_text" db:"post_text"`
}
