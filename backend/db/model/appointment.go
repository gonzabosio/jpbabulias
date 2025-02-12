package model

import "time"

type Appointment struct {
	ID       string    `json:"id"`
	ApptDate time.Time `json:"appt_date" validate:"required"`
	Subject  string    `json:"subject" validate:"required"`
	UserID   string    `json:"user_id" validate:"required"`
}

type AppointmentList struct {
	ID       string    `json:"id"`
	ApptDate time.Time `json:"appt_date"`
	Subject  string    `json:"subject"`
}
