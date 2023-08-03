package service

import (
	"context"
	"github.com/kozhamseitova/api-blog/pkg/jwttoken"
	"github.com/kozhamseitova/ustaz-hub-api/internal/config"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
	"github.com/kozhamseitova/ustaz-hub-api/internal/repository"
)

type UserService struct {
	repository repository.User
	config     *config.Config
	token      *jwttoken.JWTToken
}

func NewUserService(repository repository.User, config *config.Config, token *jwttoken.JWTToken) *UserService {
	return &UserService{
		repository: repository,
		config:     config,
		token:      token,
	}
}

func (s *UserService) CreateUser(ctx context.Context, u *entity.User) error {
	return nil
}

func (s *UserService) Login(ctx context.Context, username, password string) (string, error) {
	return "", nil
}

func (s *UserService) UpdateUser(ctx context.Context, u *entity.User) error {
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return nil
}

func (s *UserService) VerifyToken(token string) (int64, error) {
	return 0, nil
}
