package service

import (
	"context"
	"github.com/kozhamseitova/api-blog/pkg/jwttoken"
	"github.com/kozhamseitova/ustaz-hub-api/internal/config"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
	"github.com/kozhamseitova/ustaz-hub-api/internal/repository"
)

type ProductService struct {
	repository repository.Product
	config     *config.Config
	token      *jwttoken.JWTToken
}

func NewProductService(repository repository.Product, config *config.Config, token *jwttoken.JWTToken) *ProductService {
	return &ProductService{
		repository: repository,
		config:     config,
		token:      token,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, p *entity.Product) error {
	return s.repository.CreateProduct(ctx, p)
}

func (s *ProductService) GetProductById(ctx context.Context, id int64) (*entity.Product, error) {
	return s.repository.GetProductById(ctx, id)
}

func (s *ProductService) GetProductsByUserId(ctx context.Context, userId int64) ([]*entity.Product, error) {
	return s.repository.GetProductsByUserId(ctx, userId)
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	return s.repository.GetAllProducts(ctx)
}

func (s *ProductService) UpdateProduct(ctx context.Context, p *entity.Product) error {
	return s.repository.UpdateProduct(ctx, p)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	return s.repository.DeleteProduct(ctx, id)
}
