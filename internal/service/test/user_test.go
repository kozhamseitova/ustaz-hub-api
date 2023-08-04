package test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/kozhamseitova/aisha/pkg/jwttoken"
	"github.com/kozhamseitova/ustaz-hub-api/internal/config"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
	mock_repository "github.com/kozhamseitova/ustaz-hub-api/internal/repository/mock"
	service2 "github.com/kozhamseitova/ustaz-hub-api/internal/service"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
	"time"
)

func TestUserService_CreateUser(t *testing.T) {
	cfg, err := config.InitConfig("../../../config.yaml")
	require.NoError(t, err)

	jwtToken := jwttoken.New("kdjbfjvhbflkd")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUser(ctrl)

	service := service2.NewUserService(mockRepo, cfg, jwtToken)

	ctx := context.Background()

	u := &entity.User{
		ID:        1,
		Username:  "qwerty5",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:      "teacher",
		FirstName: "ff",
		LastName:  "ll",
		About:     "dkkd",
		Password:  "qwerty",
	}

	mockRepo.EXPECT().CreateUser(ctx, u).Return(nil)

	err = service.CreateUser(ctx, u)
	require.NoError(t, err)

	u2 := &entity.User{
		ID:        1,
		Username:  "qwerty5",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:      "teacher",
		FirstName: "ff",
		LastName:  "ll",
		About:     "dkkd",
		Password:  "",
	}

	mockRepo.EXPECT().CreateUser(ctx, u2).Return(fmt.Errorf("password is nil"))

	err = service.CreateUser(ctx, u2)
	require.Error(t, err)

}

func TestUserService_Login(t *testing.T) {
	cfg, err := config.InitConfig("../../../config.yaml")
	require.NoError(t, err)

	jwtToken := jwttoken.New("kdjbfjvhbflkd")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUser(ctrl)

	service := service2.NewUserService(mockRepo, cfg, jwtToken)

	ctx := context.Background()

	// Test case 1: Successful login
	username := "testuser"
	password := "testpassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)

	user := &entity.User{
		ID:       1,
		Username: username,
		Password: string(hashedPassword),
		Role:     "teacher",
	}

	mockRepo.EXPECT().GetUserByUsername(ctx, username).Return(user, nil)

	token, err := service.Login(ctx, username, password)
	log.Println(token)
	require.NoError(t, err)

	mockRepo.EXPECT().GetUserByUsername(ctx, username).Return(nil, fmt.Errorf("user not found"))

	token2, err := service.Login(ctx, username, password)
	log.Println(token2)
	require.Error(t, err)
}
