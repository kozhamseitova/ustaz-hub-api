package entity

import "time"

type User struct {
	ID           int64         `json:"id" db:"id"`
	Username     string        `json:"username" db:"username" binding:"required"`
	CreatedAt    time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" db:"updated_at"`
	Role         string        `json:"role" db:"role" binding:"required"`
	FirstName    string        `json:"first_name" db:"first_name" binding:"required"`
	LastName     string        `json:"last_name" db:"last_name" binding:"required"`
	About        string        `json:"about" db:"about"`
	Organization *Organization `json:"organization" db:"organization" binding:"required"`
	Speciality   *Speciality   `json:"speciality" db:"speciality" binding:"required"`
	Password     string        `json:"password" db:"password" binding:"required"`
}
