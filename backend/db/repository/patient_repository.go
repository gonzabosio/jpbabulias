package repository

import (
	"strconv"

	"github.com/gonzabosio/jpbabulias/db/model"
	"github.com/lib/pq"
)

type PatientRepository interface {
	ReadPatientsByUserId(userId int) (*[]model.PatientAppts, error)
	SavePatient(patient *model.InsertPatient) error
}

var _ PatientRepository = (*PostgreService)(nil)

func (p *PostgreService) ReadPatientsByUserId(userId int) (*[]model.PatientAppts, error) {
	rows, err := p.DB.Query(`SELECT * FROM patient WHERE user_id=$1`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	patientsDataMap := make(map[int]model.Patient)
	patientsIdList := []int{}
	for rows.Next() {
		var (
			patient       model.Patient
			patientID     int
			patientUserID int
		)
		if err := rows.Scan(&patientID, &patient.FirstName, &patient.LastName, &patient.PhoneNumber, &patient.Dni, &patient.HealthInsurance, &patient.Main, &patientUserID); err != nil {
			return nil, err
		}
		patientsIdList = append(patientsIdList, patientID)
		patient.ID = strconv.Itoa(patientID)
		patient.UserID = strconv.Itoa(patientUserID)
		patientsDataMap[patientID] = patient
	}

	apptRows, err := p.DB.Query(`SELECT *
		FROM appointment
		WHERE patient_id = ANY($1)
		ORDER BY appt_date`, pq.Array(patientsIdList),
	)
	if err != nil {
		return nil, err
	}
	defer apptRows.Close()
	appointmentsMap := make(map[int][]model.ExportAppointment)
	for apptRows.Next() {
		var (
			appt          model.ExportAppointment
			apptID        int
			apptPatientID int
		)
		if err := apptRows.Scan(&apptID, &appt.ApptDate, &appt.Subject, &apptPatientID); err != nil {
			return nil, err
		}
		appt.ID = strconv.Itoa(apptID)
		appt.PatientID = strconv.Itoa(apptPatientID)
		appointmentsMap[apptPatientID] = append(appointmentsMap[apptPatientID], appt)
	}
	patientsApptList := []model.PatientAppts{}
	for _, id := range patientsIdList {
		patientsApptList = append(patientsApptList, model.PatientAppts{
			Patient:      patientsDataMap[id],
			Appointments: appointmentsMap[id],
		})
	}
	return &patientsApptList, nil
}

func (p *PostgreService) SavePatient(patient *model.InsertPatient) error {
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
