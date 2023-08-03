package service

import (
	"context"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
)

type User interface {
	CreateUser(ctx context.Context, u *entity.User) error
	Login(ctx context.Context, username, password string) (string, error)
	UpdateUser(ctx context.Context, u *entity.User) error
	DeleteUser(ctx context.Context, id int64) error
	VerifyToken(token string) (int64, error)
}

type Product interface {
	CreateProduct(ctx context.Context, p *entity.Product) error
	GetProductById(ctx context.Context, id int64) (*entity.Product, error)
	GetProductsByUserId(ctx context.Context, userId int64) ([]*entity.Product, error)
	GetAllProducts(ctx context.Context) ([]*entity.Product, error)
	UpdateProduct(ctx context.Context, p *entity.Product) error
	DeleteProduct(ctx context.Context, id int64) error
}

type Post interface {
	CreatePost(ctx context.Context, p *entity.Post) error
	GetPostById(ctx context.Context, id int64) (*entity.Post, error)
	GetPostsByUserId(ctx context.Context, userId int64) ([]*entity.Post, error)
	GetAllPosts(ctx context.Context) ([]*entity.Post, error)
	UpdatePost(ctx context.Context, p *entity.Post) error
	DeletePost(ctx context.Context, id int64) error
}

type Comment interface {
	CreateComment(ctx context.Context, c *entity.Comment) error
	GetAllProductsComments(ctx context.Context, productId int64) ([]*entity.Comment, error)
	GetAllPostsComments(ctx context.Context, postId int64) ([]*entity.Comment, error)
	GetCommentsByParentId(ctx context.Context, parentId int64) ([]*entity.Comment, error)
}