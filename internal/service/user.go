package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/kozhamseitova/aisha/pkg/jwttoken"
	"github.com/kozhamseitova/aisha/pkg/util"
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
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = s.repository.CreateUser(ctx, u)

	if err != nil {
		return fmt.Errorf("create user err: %w", err)
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.repository.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found")
		}
		return "", fmt.Errorf("get user err: %w", err)
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return "", fmt.Errorf("incorrect password err: %w", err)
	}

	token, err := s.token.CreateToken(user.ID, user.Role, s.config.Token.TimeToLive)

	if err != nil {
		return "", fmt.Errorf("create token err: %w", err)
	}

	return token, nil
}

func (s *UserService) UpdateUser(ctx context.Context, u *entity.User) error {

	err := s.repository.UpdateUser(ctx, u)
	if err != nil {
		return fmt.Errorf("update user err: %w", err)
	}
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	err := s.repository.DeleteUser(ctx, id)

	if err != nil {
		return fmt.Errorf("delete user err: %w", err)
	}
	return nil
}

func (s *UserService) VerifyToken(token string) (int64, string, error) {
	payload, err := s.token.ValidateToken(token)
	if err != nil {
		return 0, "", fmt.Errorf("validate token err: %w", err)
	}

	return payload.UserId, payload.UserRole, nil
}

func (s *UserService) CheckPermission(userId, objUserId int64, userRole, action string) bool {
	switch userRole {
	case "admin":
		return true
	case "moderator":
		return action == "delete" || action == "update" || action == "get"
	case "teacher":
		if userId != objUserId {
			return false
		}
		return action == "update" || action == "get" || action == "create"
	default:
		return false
	}
}
