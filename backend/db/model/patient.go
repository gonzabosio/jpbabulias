package model

type Patient struct {
	ID              string `json:"id"`
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	PhoneNumber     int    `json:"phone_number" validate:"required"`
	Dni             int    `json:"dni" validate:"required"`
	HealthInsurance string `json:"health_insurance" validate:"required"`
	Main            bool   `json:"main" validate:"required"`
	UserID          string `json:"user_id" validate:"required"`
}
