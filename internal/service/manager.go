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

func NewService(repository repository.Repository, config config.Config, token *jwttoken.JWTToken) *Service {
	return &Service{
		Authorization: NewAuthService(repository, config, token),
		User:          NewUserService(repository, config, token),
		Product:       NewProductService(repository, config, token),
		Post:          NewPostService(repository, config, token),
		Comment:       NewCommentService(repository, config, token),
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
