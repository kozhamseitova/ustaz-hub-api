package entity

import "time"

type User struct {
	ID           int64         `json:"id" db:"id"`
	Username     string        `json:"username" db:"username" binding:"required"`
	CreatedAt    time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" db:"updated_at"`
	Role         string        `json:"role" db:"role"`
	FirstName    string        `json:"first_name" db:"first_name"`
	LastName     string        `json:"last_name" db:"last_name"`
	About        string        `json:"about" db:"about"`
	Organization *Organization `json:"organization" db:"organization"`
	Speciality   *Speciality   `json:"speciality" db:"speciality"`
	Password     string        `json:"password" db:"password"`
}
