package pgrepo

import "github.com/jackc/pgx/v4/pgxpool"

type AuthPostgres struct {
	Pool *pgxpool.Pool
}

func NewAuthPostgres(pool *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{
		Pool: pool,
	}
}
