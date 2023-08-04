package service

import (
	"context"
	"fmt"
	"github.com/kozhamseitova/aisha/pkg/jwttoken"
	"github.com/kozhamseitova/ustaz-hub-api/internal/config"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
	"github.com/kozhamseitova/ustaz-hub-api/internal/repository"
)

type PostService struct {
	repository repository.Post
	config     *config.Config
	token      *jwttoken.JWTToken
}

func NewPostService(repository repository.Post, config *config.Config, token *jwttoken.JWTToken) *PostService {
	return &PostService{
		repository: repository,
		config:     config,
		token:      token,
	}
}

func (s *PostService) CreatePost(ctx context.Context, p *entity.Post) (int64, error) {
	return s.repository.CreatePost(ctx, p)
}

func (s *PostService) GetPostById(ctx context.Context, id int64) (*entity.Post, error) {
	return s.repository.GetPostById(ctx, id)
}

func (s *PostService) GetPostsByUserId(ctx context.Context, userId int64) ([]entity.Post, error) {
	return s.repository.GetPostsByUserId(ctx, userId)
}

func (s *PostService) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	return s.repository.GetAllPosts(ctx)
}

func (s *PostService) UpdatePost(ctx context.Context, p *entity.Post) error {
	return s.repository.UpdatePost(ctx, p)
}

func (s *PostService) DeletePost(ctx context.Context, id int64) error {
	err := s.repository.DeletePost(ctx, id)

	if err != nil {
		return fmt.Errorf("delete post err: %w", err)
	}

	return nil
}
