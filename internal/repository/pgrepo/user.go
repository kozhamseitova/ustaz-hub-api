package pgrepo

import "github.com/jackc/pgx/v4/pgxpool"

type UserPostgres struct {
	Pool *pgxpool.Pool
}

func NewUserPostgres(pool *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{
		Pool: pool,
	}
}
