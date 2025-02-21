package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
	Admin     bool      `json:"admin"`
}

type UserSignUp struct {
	ID              string    `json:"id"`
	FirstName       string    `json:"first_name" validate:"required"`
	LastName        string    `json:"last_name" validate:"required"`
	Email           string    `json:"email" validate:"required,email"`
	Password        string    `json:"password" validate:"required"`
	BirthDate       time.Time `json:"birth_date" validate:"required"`
	Admin           bool      `json:"admin"`
	PhoneNumber     string    `json:"phone_number" validate:"required"`
	Dni             int       `json:"dni" validate:"required"`
	HealthInsurance string    `json:"health_insurance" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Admin bool   `json:"admin"`
}
