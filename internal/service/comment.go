package service

import (
	"context"
	"github.com/kozhamseitova/api-blog/pkg/jwttoken"
	"github.com/kozhamseitova/ustaz-hub-api/internal/config"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
	"github.com/kozhamseitova/ustaz-hub-api/internal/repository"
)

type CommentService struct {
	repository repository.Comment
	config     *config.Config
	token      *jwttoken.JWTToken
}

func NewCommentService(repository repository.Comment, config *config.Config, token *jwttoken.JWTToken) *CommentService {
	return &CommentService{
		repository: repository,
		config:     config,
		token:      token,
	}
}

func (s *CommentService) CreateComment(ctx context.Context, c *entity.Comment) error {
	return s.repository.CreateComment(ctx, c)
}

func (s *CommentService) GetAllProductsComments(ctx context.Context, productId int64) ([]*entity.Comment, error) {
	return s.repository.GetAllProductsComments(ctx, productId)
}

func (s *CommentService) GetAllPostsComments(ctx context.Context, postId int64) ([]*entity.Comment, error) {
	return s.repository.GetAllPostsComments(ctx, postId)
}

func (s *CommentService) GetCommentsByParentId(ctx context.Context, parentId int64) ([]*entity.Comment, error) {
	return s.repository.GetCommentsByParentId(ctx, parentId)
}
