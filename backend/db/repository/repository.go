package repository

import "database/sql"

type Service interface {
	AppointmentRepository
}

type PostgreService struct {
	DB *sql.DB
}
