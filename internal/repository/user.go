package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kozhamseitova/ustaz-hub-api/internal/entity"
	"strings"
)

type UserPostgres struct {
	Pool *pgxpool.Pool
}

func NewUserPostgres(pool *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{
		Pool: pool,
	}
}

func (p *UserPostgres) CreateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`INSERT INTO %s (
												username,
												password,
												role,
												first_name,
												last_name,
												about,
												organization_id,
												speciality_id
											)
									VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, usersTable)

	_, err := p.Pool.Exec(ctx, query, u.Username, u.Password, u.Role, u.FirstName, u.LastName, u.About, u.Organization.ID, u.Speciality.ID)

	if err != nil {
		return err
	}

	return nil
}

func (p *UserPostgres) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := &entity.User{}

	query := fmt.Sprintf(`
						SELECT u.id, u.username, u.password, u.created_at, u.updated_at, u.role, u.first_name, u.last_name, u.about,
						   o.id as "organization.id", o.name as "organization.name", o.location as "organization.location",
						   s.id as "speciality.id", s.name as "speciality.name"
						FROM %s u
						JOIN %s o on o.id = u.organization_id
						JOIN %s s on s.id = u.speciality_id
						WHERE
							u.username = $1
						LIMIT 1`, usersTable, organizationsTable, specialitiesTable)

	err := pgxscan.Get(ctx, p.Pool, user, query, strings.TrimSpace(username))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *UserPostgres) UpdateUser(ctx context.Context, u *entity.User) error {
	return nil
}

func (p *UserPostgres) DeleteUser(ctx context.Context, id int64) error {
	return nil
}
