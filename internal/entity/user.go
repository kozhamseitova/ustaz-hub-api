package entity

import "time"

type User struct {
	ID           int64         `json:"id" db:"id""`
	Email        string        `json:"email" db:"email""`
	CreatedAt    time.Duration `json:"created_at" db:"created_at""`
	UpdatedAt    time.Duration `json:"updated_at" db:"updated_at""`
	Role         string        `json:"role" db:"role""`
	FirstName    string        `json:"first_name" db:"first_name""`
	LastName     string        `json:"last_name" db:"last_name""`
	About        string        `json:"about" db:"about""`
	Organization Organization  `json:"organization" db:"organization""`
	Speciality   Speciality    `json:"speciality" db:"speciality"`
}
