package service

import (
	"github.com/kozhamseitova/api-blog/pkg/jwttoken"
	"github.com/kozhamseitova/ustaz-hub-api/internal/config"
	"github.com/kozhamseitova/ustaz-hub-api/internal/repository"
)

type AuthService struct {
	repository repository.Repository
	config     config.Config
	token      *jwttoken.JWTToken
}

func NewAuthService(repository repository.Repository, config config.Config, token *jwttoken.JWTToken) *AuthService {
	return &AuthService{
		repository: repository,
		config:     config,
		token:      token,
	}
}
