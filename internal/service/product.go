package service

import (
	"github.com/kozhamseitova/api-blog/pkg/jwttoken"
	"github.com/kozhamseitova/ustaz-hub-api/internal/config"
	"github.com/kozhamseitova/ustaz-hub-api/internal/repository"
)

type ProductService struct {
	repository repository.Product
	config     *config.Config
	token      *jwttoken.JWTToken
}

func NewProductService(repository repository.Product, config *config.Config, token *jwttoken.JWTToken) *AuthService {
	return &AuthService{
		repository: repository,
		config:     config,
		token:      token,
	}
}
