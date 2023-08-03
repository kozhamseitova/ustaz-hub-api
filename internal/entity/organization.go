package entity

type Organization struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Location string `json:"location" db:"location"`
}
