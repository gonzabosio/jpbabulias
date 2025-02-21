package model

import "time"

type Appointment struct {
	ID        int       `json:"id"`
	ApptDate  time.Time `json:"appt_date" validate:"required"`
	Subject   string    `json:"subject" validate:"required"`
	PatientID int       `json:"patient_id" validate:"required"`
}

type ExportAppointment struct {
	ID        string    `json:"id"`
	ApptDate  time.Time `json:"appt_date"`
	Subject   string    `json:"subject"`
	PatientID string    `json:"patient_id"`
}

type InsertAppointment struct {
	ID        string    `json:"id"`
	ApptDate  time.Time `json:"appt_date" validate:"required"`
	Subject   string    `json:"subject" validate:"required"`
	PatientID string    `json:"patient_id" validate:"required"`
	UserEmail string    `json:"email" validate:"required"`
	FullName  string    `json:"full_name"`
}

type AppointmentList struct {
	ID              int       `json:"id"`
	ApptDate        time.Time `json:"appt_date"`
	Subject         string    `json:"subject"`
	PatientID       int       `json:"patient_id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Dni             int       `json:"dni"`
	HealthInsurance string    `json:"health_insurance"`
	PhoneNumber     int       `json:"phone_number"`
	UserID          int       `json:"user_id"`
	Main            bool      `json:"main"`
}
