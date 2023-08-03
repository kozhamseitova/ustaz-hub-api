package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
)

type UserPostgres struct {
	Pool *pgxpool.Pool
}

func NewUserPostgres(pool *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{
		Pool: pool,
	}
}

func (p *UserPostgres) CreateUser(ctx context.Context, u *entity.User) error {
	return nil
}

func (p *UserPostgres) GetUserByUsername(ctx context.Context, username, password string) (string, error) {
	return "", nil
}

func (p *UserPostgres) UpdateUser(ctx context.Context, u *entity.User) error {
	return nil
}

func (p *UserPostgres) DeleteUser(ctx context.Context, id int64) error {
	return nil
}
