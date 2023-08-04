package service

import (
	"context"
	"fmt"
	"github.com/kozhamseitova/aisha/pkg/jwttoken"
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

func (s *ProductService) CreateProduct(ctx context.Context, p *entity.Product) (int64, error) {
	id, err := s.repository.CreateProduct(ctx, p)

	if err != nil {
		return 0, fmt.Errorf("create product err: %w", err)
	}

	return id, nil
}

func (s *ProductService) GetProductById(ctx context.Context, id int64) (*entity.Product, error) {
	product, err := s.repository.GetProductById(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("get product by id err: %w", err)
	}
	return product, nil
}

func (s *ProductService) GetProductsByUserId(ctx context.Context, userId int64) ([]entity.Product, error) {
	products, err := s.repository.GetProductsByUserId(ctx, userId)

	if err != nil {
		return nil, fmt.Errorf("get products by user id err: %w", err)
	}
	return products, nil
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	products, err := s.repository.GetAllProducts(ctx)

	if err != nil {
		return nil, fmt.Errorf("get all products err: %w", err)
	}
	return products, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, p *entity.Product) error {
	err := s.repository.UpdateProduct(ctx, p)
	if err != nil {
		return fmt.Errorf("update product err: %w", err)
	}
	return nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	err := s.repository.DeleteProduct(ctx, id)

	if err != nil {
		return fmt.Errorf("delete product err: %w", err)
	}
	return nil
}
