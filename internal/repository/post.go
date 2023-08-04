package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
)

type PostPostgres struct {
	Pool *pgxpool.Pool
}

func NewPostPostgres(pool *pgxpool.Pool) *PostPostgres {
	return &PostPostgres{
		Pool: pool,
	}
}

func (p *PostPostgres) CreatePost(ctx context.Context, post *entity.Post) (int64, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			user_id,
			post_text
		)
		VALUES ($1, $2)
		RETURNING id`, postsTable)

	var postID int64
	err := pgxscan.Get(ctx, p.Pool, &postID, query, post.User.ID, post.PostText)

	if err != nil {
		return 0, err
	}

	return postID, nil
}

func (p *PostPostgres) GetPostById(ctx context.Context, id int64) (*entity.Post, error) {
	var post = new(entity.Post)

	query := fmt.Sprintf(`
						SELECT p.id, p.post_text,
						       u.id as "user.id", u.username as "user.username" 
						FROM %s p
						JOIN %s u on u.id = p.user_id                                      
						WHERE p.id = $1
						LIMIT 1`, postsTable, usersTable)

	err := pgxscan.Get(ctx, p.Pool, post, query, id)

	if err != nil {
		return post, err
	}

	return post, nil
}

func (p *PostPostgres) GetPostsByUserId(ctx context.Context, userId int64) ([]entity.Post, error) {
	var posts []entity.Post

	query := fmt.Sprintf(`
						SELECT p.id, p.post_text,
						       u.id as "user.id", u.username as "user.username" 
						FROM %s p
						JOIN %s u on u.id = p.user_id                                      
						WHERE u.id = $1`, postsTable, usersTable)

	err := pgxscan.Select(ctx, p.Pool, &posts, query, userId)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostPostgres) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	var posts []entity.Post

	query := fmt.Sprintf(`
						SELECT p.id, p.post_text,
						       u.id as "user.id", u.username as "user.username" 
						FROM %s p
						JOIN %s u on u.id = p.user_id`, postsTable, usersTable)

	err := pgxscan.Select(ctx, p.Pool, &posts, query)

	if err != nil {
		return posts, err
	}

	return posts, nil
}

func (p *PostPostgres) UpdatePost(ctx context.Context, post *entity.Post) error {
	query := fmt.Sprintf(
		`UPDATE %s SET
								post_text = $1,
							WHERE id = $2`, postsTable)

	_, err := p.Pool.Exec(ctx, query,
		post.PostText,
		post.ID,
	)

	return err
}

func (p *PostPostgres) DeletePost(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, postsTable)

	_, err := p.Pool.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
