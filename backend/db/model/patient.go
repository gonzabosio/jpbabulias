package model

type Patient struct {
	ID              string `json:"id"`
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	PhoneNumber     string `json:"phone_number" validate:"required"`
	Dni             string `json:"dni" validate:"required"`
	HealthInsurance string `json:"health_insurance" validate:"required"`
	Main            bool   `json:"main"`
	UserID          string `json:"user_id" validate:"required"`
}

type PatientAppts struct {
	Patient      Patient             `json:"patient"`
	Appointments []ExportAppointment `json:"appointments"`
}

type InsertPatient struct {
	ID              string `json:"id"`
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	PhoneNumber     string `json:"phone_number" validate:"required"`
	Dni             int    `json:"dni" validate:"required"`
	HealthInsurance string `json:"health_insurance" validate:"required"`
	Main            bool   `json:"main"`
	UserID          string `json:"user_id" validate:"required"`
}
