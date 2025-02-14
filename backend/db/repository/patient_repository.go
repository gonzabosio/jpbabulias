package repository

import "github.com/gonzabosio/jpbabulias/db/model"

type PatientRepository interface {
	ReadPatientsByUserId(userId int) (*[]model.Patient, error)
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
