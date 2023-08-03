package entity

type Organization struct {
	ID       int64 `json:"id" db:"id""`
	Name     int64 `json:"name" db:"name""`
	Location int64 `json:"location" db:"location""`
}
