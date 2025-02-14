package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	BirthDate time.Time `json:"birth_date" validate:"required"`
	Admin     bool      `json:"admin" validate:"required"`
}
