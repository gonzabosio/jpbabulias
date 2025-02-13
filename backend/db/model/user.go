package model

import "time"

type User struct {
	ID              string    `json:"id"`
	FirstName       string    `json:"first_name" validate:"required"`
	LastName        string    `json:"last_name" validate:"required"`
	BirthDate       time.Time `json:"birth_date" validate:"required"`
	Email           string    `json:"email" validate:"required"`
	Password        string    `json:"password" validate:"required"`
	PhoneNumber     int       `json:"phone_number" validate:"required"`
	Dni             int       `json:"dni" validate:"required"`
	HealthInsurance string    `json:"health_insurance" validate:"required"`
	Admin           bool      `json:"admin" validate:"required"`
}
