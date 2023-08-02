package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kozhamseitova/ustaz-hub-api/internal/repository/pgrepo"
)

type Repository struct {
	Authorization
	User
	Product
	Post
	Comment
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: pgrepo.NewAuthPostgres(pool),
		User:          pgrepo.NewUserPostgres(pool),
		Product:       pgrepo.NewProductPostgres(pool),
		Post:          pgrepo.NewPostPostgres(pool),
		Comment:       pgrepo.NewCommentPostgres(pool),
	}
}

type Authorization interface {
}

type User interface {
}

type Product interface {
}

type Post interface {
}

type Comment interface {
}
