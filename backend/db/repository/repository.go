package repository

import "database/sql"

type Service interface {
	AppointmentRepository
	PatientRepository
}

type PostgreService struct {
	DB *sql.DB
}
