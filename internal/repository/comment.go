package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
)

type CommentPostgres struct {
	Pool *pgxpool.Pool
}

func NewCommentPostgres(pool *pgxpool.Pool) *CommentPostgres {
	return &CommentPostgres{
		Pool: pool,
	}
}

func (p *CommentPostgres) CreateComment(ctx context.Context, c *entity.Comment) error {
	return nil
}

func (p *CommentPostgres) GetAllProductsComments(ctx context.Context, productId int64) ([]*entity.Comment, error) {
	return nil, nil
}

func (p *CommentPostgres) GetAllPostsComments(ctx context.Context, postId int64) ([]*entity.Comment, error) {
	return nil, nil
}

func (p *CommentPostgres) GetCommentsByParentId(ctx context.Context, parentId int64) ([]*entity.Comment, error) {
	return nil, nil
}
