package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	usersTable         = "users"
	productsTable      = "products"
	postsTable         = "posts"
	commentsTable      = "comments"
	specialitiesTable  = "specialities"
	organizationsTable = "organizations"
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
