package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
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

func (p *ProductPostgres) CreateProduct(ctx context.Context, product *entity.Product) (int64, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			user_id,
			title,
			description,
			file_path
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id`, productsTable)

	var productID int64
	err := pgxscan.Get(ctx, p.Pool, &productID, query, product.User.ID, product.Title, product.Description, product.FilePath)

	if err != nil {
		return 0, err
	}

	return productID, nil
}

func (p *ProductPostgres) GetProductById(ctx context.Context, id int64) (*entity.Product, error) {
	var product = new(entity.Product)

	query := fmt.Sprintf(`
						SELECT p.id, p.title, p.description, p.file_path, p.uploaded_at,
						       u.id as "user.id", u.username as "user.username" 
						FROM %s p
						JOIN %s u on u.id = p.user_id                                      
						WHERE p.id = $1
						LIMIT 1`, productsTable, usersTable)

	err := pgxscan.Get(ctx, p.Pool, product, query, id)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductPostgres) GetProductsByUserId(ctx context.Context, userId int64) ([]entity.Product, error) {
	var products []entity.Product

	query := fmt.Sprintf(`
						SELECT p.id, p.title, p.description, p.file_path, p.uploaded_at,
						       u.id as "user.id", u.username as "user.username" 
						FROM %s p
						JOIN %s u on u.id = p.user_id                                      
						WHERE u.id = $1`, productsTable, usersTable)

	err := pgxscan.Select(ctx, p.Pool, &products, query, userId)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductPostgres) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	var products []entity.Product

	query := fmt.Sprintf(`
						SELECT p.id, p.title, p.description, p.file_path, p.uploaded_at,
						       u.id as "user.id", u.username as "user.username" 
						FROM %s p
						JOIN %s u on u.id = p.user_id`, productsTable, usersTable)

	err := pgxscan.Select(ctx, p.Pool, &products, query)

	if err != nil {
		return products, err
	}

	return products, nil
}

func (p *ProductPostgres) UpdateProduct(ctx context.Context, product *entity.Product) error {
	query := fmt.Sprintf(
		`UPDATE %s SET
								title = $1,
								description = $2,
								file_path = $3
							WHERE id = $4`, productsTable)

	_, err := p.Pool.Exec(ctx, query,
		product.Title,
		product.Description,
		product.FilePath,
		product.ID,
	)

	return err
}

func (p *ProductPostgres) DeleteProduct(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, productsTable)

	_, err := p.Pool.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
