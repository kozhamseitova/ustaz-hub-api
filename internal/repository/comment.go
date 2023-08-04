package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
)

type CommentPostgres struct {
	Pool *pgxpool.Pool
}

func NewCommentPostgres(pool *pgxpool.Pool) *CommentPostgres {
	return &CommentPostgres{
		Pool: pool,
	}
}

func (p *CommentPostgres) CreateComment(ctx context.Context, c *entity.CreateComment) (int64, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (parent_id, comment_text, author_id, is_review, product_id, grade, post_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`, commentsTable)

	var commentID int64
	err := p.Pool.QueryRow(ctx, query, c.ParentID, c.CommentText, c.AuthorID, c.IsReview, c.ProductID, c.Grade, c.PostID).Scan(&commentID)
	if err != nil {
		return 0, err
	}

	return commentID, nil
}

func (p *CommentPostgres) GetAllProductsComments(ctx context.Context, productId int64) ([]entity.Comment, error) {
	query := fmt.Sprintf(`
							WITH RECURSIVE comments_recursive AS (
								SELECT
									id,
									parent_id,
									comment_text,
									grade,
									created_at
								FROM comments
								WHERE parent_id IS NULL
								  AND is_review = true
								  AND product_id is not null
							
								UNION ALL
							
								SELECT
									c.id,
									c.parent_id,
									c.comment_text,
									c.grade,
									c.created_at
								FROM comments c
										 JOIN comments_recursive cr ON c.parent_id = cr.id
							)
							SELECT
								cr.id,
								coalesce(cr.parent_id, 0) as parent_id,
								cr.comment_text,
								cr.grade,
								cr.created_at::timestamptz,
								(
									SELECT jsonb_agg(jsonb_build_object(
											'id', sub_c.id,
											'parent_id', sub_c.parent_id,
											'comment_text', sub_c.comment_text,
											'grade', sub_c.grade,
											'created_at', sub_c.created_at::timestamptz,
											'comments', (
												SELECT jsonb_agg(jsonb_build_object(
														'id', sub_sub_c.id,
														'parent_id', sub_sub_c.parent_id,
														'comment_text', sub_sub_c.comment_text,
														'grade', sub_sub_c.grade,
														'created_at', sub_sub_c.created_at::timestamptz
													))
												FROM comments sub_sub_c
												WHERE sub_c.id = sub_sub_c.parent_id
											)
										))
									FROM comments sub_c
									WHERE cr.id = sub_c.parent_id
								) AS comments
							FROM comments_recursive cr
							WHERE cr.parent_id IS NULL
							GROUP BY cr.id, cr.parent_id, cr.comment_text, cr.grade, cr.created_at;
					`)

	var comments []entity.Comment
	err := pgxscan.Select(ctx, p.Pool, &comments, query)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (p *CommentPostgres) GetAllPostsComments(ctx context.Context, postId int64) ([]entity.Comment, error) {
	query := fmt.Sprintf(`
						WITH RECURSIVE comments_recursive AS (
						SELECT
							id,
							parent_id,
							comment_text,
							created_at
						FROM comments
						WHERE parent_id IS NULL
						  AND post_id is not null
					
						UNION ALL
					
						SELECT
							c.id,
							c.parent_id,
							c.comment_text,
							c.created_at
						FROM comments c
								 JOIN comments_recursive cr ON c.parent_id = cr.id
					)
					SELECT
						cr.id,
						coalesce(cr.parent_id, 0) as parent_id,
						cr.comment_text,
						cr.created_at::timestamptz,
						(
							SELECT jsonb_agg(jsonb_build_object(
									'id', sub_c.id,
									'parent_id', sub_c.parent_id,
									'comment_text', sub_c.comment_text,
									'created_at', sub_c.created_at::timestamptz,
									'comments', (
										SELECT jsonb_agg(jsonb_build_object(
												'id', sub_sub_c.id,
												'parent_id', sub_sub_c.parent_id,
												'comment_text', sub_sub_c.comment_text,
												'created_at', sub_sub_c.created_at::timestamptz
											))
										FROM comments sub_sub_c
										WHERE sub_c.id = sub_sub_c.parent_id
									)
								))
							FROM comments sub_c
							WHERE cr.id = sub_c.parent_id
						) AS comments
					FROM comments_recursive cr
					WHERE cr.parent_id IS NULL
					GROUP BY cr.id, cr.parent_id, cr.comment_text, cr.created_at;
					`)

	var comments []entity.Comment
	err := pgxscan.Select(ctx, p.Pool, &comments, query)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (p *CommentPostgres) GetCommentsByParentId(ctx context.Context, parentId int64) ([]entity.Comment, error) {
	return nil, nil
}
