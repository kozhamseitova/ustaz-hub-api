package pgrepo

import "github.com/jackc/pgx/v4/pgxpool"

type PostPostgres struct {
	Pool *pgxpool.Pool
}

func NewPostPostgres(pool *pgxpool.Pool) *PostPostgres {
	return &PostPostgres{
		Pool: pool,
	}
}
