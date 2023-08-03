package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
)

type ProductPostgres struct {
	Pool *pgxpool.Pool
}

func NewProductPostgres(pool *pgxpool.Pool) *ProductPostgres {
	return &ProductPostgres{
		Pool: pool,
	}
}

func (p *ProductPostgres) CreateProduct(ctx context.Context, product *entity.Product) error {
	return nil
}

func (p *ProductPostgres) GetProductById(ctx context.Context, id int64) (*entity.Product, error) {
	return nil, nil
}

func (p *ProductPostgres) GetProductsByUserId(ctx context.Context, userId int64) ([]*entity.Product, error) {
	return nil, nil
}

func (p *ProductPostgres) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	return nil, nil
}

func (p *ProductPostgres) UpdateProduct(ctx context.Context, product *entity.Product) error {
	return nil
}

func (p *ProductPostgres) DeleteProduct(ctx context.Context, id int64) error {
	return nil
}
