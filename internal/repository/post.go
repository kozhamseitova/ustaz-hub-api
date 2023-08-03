package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
)

type PostPostgres struct {
	Pool *pgxpool.Pool
}

func NewPostPostgres(pool *pgxpool.Pool) *PostPostgres {
	return &PostPostgres{
		Pool: pool,
	}
}

func (p *PostPostgres) CreatePost(ctx context.Context, post *entity.Post) error {
	return nil
}

func (p *PostPostgres) GetPostById(ctx context.Context, id int64) (*entity.Post, error) {
	return nil, nil
}

func (p *PostPostgres) GetPostsByUserId(ctx context.Context, userId int64) ([]*entity.Post, error) {
	return nil, nil
}

func (p *PostPostgres) GetAllPosts(ctx context.Context) ([]*entity.Post, error) {
	return nil, nil
}

func (p *PostPostgres) UpdatePost(ctx context.Context, post *entity.Post) error {
	return nil
}

func (p *PostPostgres) DeletePost(ctx context.Context, id int64) error {
	return nil
}
