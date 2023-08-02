package service

import (
	"github.com/kozhamseitova/api-blog/pkg/jwttoken"
	"github.com/kozhamseitova/ustaz-hub-api/internal/config"
	"github.com/kozhamseitova/ustaz-hub-api/internal/repository"
)

type Service struct {
	Authorization
	User
	Product
	Post
	Comment
}

func NewService(repository *repository.Repository, config *config.Config, token *jwttoken.JWTToken) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization, config, token),
		User:          NewUserService(repository.User, config, token),
		Product:       NewProductService(repository.User, config, token),
		Post:          NewPostService(repository.Post, config, token),
		Comment:       NewCommentService(repository.Comment, config, token),
	}
}

type Authorization interface {
}

type User interface {
}

type Product interface {
}

type Post interface {
}

type Comment interface {
}
