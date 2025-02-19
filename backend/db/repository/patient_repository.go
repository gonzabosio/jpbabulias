package repository

import (
	"strconv"

	"github.com/gonzabosio/jpbabulias/db/model"
)

type PatientRepository interface {
	ReadPatientsByUserId(userId int) (*[]model.Patient, error)
	SavePatient(patient *model.Patient) error
}

var _ PatientRepository = (*PostgreService)(nil)

func (p *PostgreService) ReadPatientsByUserId(userId int) (*[]model.Patient, error) {
	rows, err := p.DB.Query(`SELECT * FROM patient WHERE user_id=$1`, userId)
	if err != nil {
		return nil, err
	}
	patientList := []model.Patient{}
	for rows.Next() {
		var patient model.Patient
		if err := rows.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.PhoneNumber, &patient.Dni, &patient.HealthInsurance, &patient.Main, &patient.UserID); err != nil {
			return nil, err
		}
		patientList = append(patientList, patient)
	}
	return &patientList, nil
}

func (p *PostgreService) SavePatient(patient *model.Patient) error {
	row := p.DB.QueryRow(`INSERT INTO patient (first_name, last_name, phone_number, dni, health_insurance, main, user_id) 
	VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		patient.FirstName, patient.LastName, patient.PhoneNumber, patient.Dni, patient.HealthInsurance, false, patient.UserID)

	var patientId int
	if err := row.Scan(&patientId); err != nil {
		return err
	}
	patientIdStr := strconv.Itoa(patientId)
	patient.ID = patientIdStr
	return nil
}
