package pgrepo

import "github.com/jackc/pgx/v4/pgxpool"

type CommentPostgres struct {
	Pool *pgxpool.Pool
}

func NewCommentPostgres(pool *pgxpool.Pool) *CommentPostgres {
	return &CommentPostgres{
		Pool: pool,
	}
}
