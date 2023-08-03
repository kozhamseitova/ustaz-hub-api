package service

import (
	"github.com/kozhamseitova/aisha/pkg/jwttoken"
	"github.com/kozhamseitova/ustaz-hub-api/internal/config"
	"github.com/kozhamseitova/ustaz-hub-api/internal/repository"
)

type Service struct {
	User
	Product
	Post
	Comment
}

func NewService(repository *repository.Repository, config *config.Config, token *jwttoken.JWTToken) *Service {
	return &Service{
		User:    NewUserService(repository.User, config, token),
		Product: NewProductService(repository.Product, config, token),
		Post:    NewPostService(repository.Post, config, token),
		Comment: NewCommentService(repository.Comment, config, token),
	}
}
