package pgrepo

import "github.com/jackc/pgx/v4/pgxpool"

type ProductPostgres struct {
	Pool *pgxpool.Pool
}

func NewProductPostgres(pool *pgxpool.Pool) *ProductPostgres {
	return &ProductPostgres{
		Pool: pool,
	}
}
