package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	User
	Product
	Post
	Comment
}

func NewPostgresRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		User:    NewUserPostgres(pool),
		Product: NewProductPostgres(pool),
		Post:    NewPostPostgres(pool),
		Comment: NewCommentPostgres(pool),
	}
}
