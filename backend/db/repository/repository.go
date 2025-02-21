package repository

import "database/sql"

type Service interface {
	AppointmentRepository
	PatientRepository
	UserRepository
}

type PostgreService struct {
	DB *sql.DB
}
