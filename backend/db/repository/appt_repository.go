package repository

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/gonzabosio/jpbabulias/db/model"
)

type AppointmentRepository interface {
	SaveAppointment(appt *model.InsertAppointment) error
	ReadAppointmentsByUserId(userIdStr string) (*[]model.AppointmentList, error)
	ReadAppointmentsByDay(day string) (*[]string, error)
	ReadFullyBookedDates() (*[]string, error)
	DeleteAppointment(apptId string) error
}

// var _ AppointmentRepository = (*PostgreService)(nil)

func (p *PostgreService) SaveAppointment(appt *model.InsertAppointment) error {
	var apptID int
	patientIdNum, err := strconv.ParseInt(appt.PatientID, 10, 64)
	if err != nil {
		return err
	}
	err = p.DB.QueryRow(`INSERT INTO appointment(appt_date, subject, patient_id) VALUES($1, $2, $3) RETURNING id`, appt.ApptDate, appt.Subject, patientIdNum).Scan(&apptID)
	if err != nil {
		return err
	}
	appt.ID = strconv.Itoa(apptID)
	return nil
}

func (p *PostgreService) ReadAppointmentsByUserId(userIdStr string) (*[]model.AppointmentList, error) {
	appts := []model.AppointmentList{}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return nil, err
	}
	rows, err := p.DB.Query(`SELECT a.id, a.appt_date, a.subject, p.* FROM appointment AS a JOIN patient AS p ON p.id=a.patient_id WHERE user_id=$1`, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var appt model.AppointmentList
		if err := rows.Scan(&appt.ID, &appt.ApptDate, &appt.Subject, &appt.PatientID, &appt.FirstName, &appt.LastName, &appt.PhoneNumber, &appt.Dni, &appt.HealthInsurance, &appt.Main, &appt.UserID); err != nil {
			return nil, err
		}
		appts = append(appts, appt)
	}
	return &appts, nil
}

func (p *PostgreService) ReadAppointmentsByDay(day string) (*[]string, error) {
	parsedDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		return nil, err
	}
	start := parsedDate.Format("2006-01-02 15:04:05")
	end := parsedDate.Add(24 * time.Hour).Format("2006-01-02 15:04:05")
	rows, err := p.DB.Query(`SELECT appt_date FROM appointment WHERE appt_date >= $1 and appt_date <= $2`, start, end)
	if err != nil {
		return nil, err
	}
	var apptList []time.Time
	hoursList := []string{}
	for rows.Next() {
		var appt time.Time
		if err := rows.Scan(&appt); err != nil {
			if err == sql.ErrNoRows {
				return &hoursList, nil
			} else {
				return nil, err
			}
		}
		apptList = append(apptList, appt)
	}
	for _, it := range apptList {
		hm := it.Format("15:04")
		hoursList = append(hoursList, hm)
	}
	return &hoursList, nil
}

func (p *PostgreService) ReadFullyBookedDates() (*[]string, error) {
	rows, err := p.DB.Query(
		`WITH fully_booked_dates AS (
		SELECT 
			DATE(appt_date) AS date,  --extract date
			COUNT(*) AS booked_slots
		FROM appointment
		WHERE appt_date BETWEEN CURRENT_DATE 
			AND (CURRENT_DATE + INTERVAL '4 months')
		GROUP BY DATE(appt_date)
		HAVING COUNT(*) = 16  --max slots
		)
		SELECT date FROM fully_booked_dates;`,
	)
	if err != nil {
		return nil, err
	}
	fullyBookedDates := []string{}
	for rows.Next() {
		var date string
		if err := rows.Scan(&date); err != nil {
			if err == sql.ErrNoRows {
				return &fullyBookedDates, nil
			} else {
				return nil, err
			}
		}
		fullyBookedDates = append(fullyBookedDates, date)
	}
	return &fullyBookedDates, nil
}

func (p *PostgreService) DeleteAppointment(apptIdStr string) error {
	apptId, err := strconv.ParseInt(apptIdStr, 10, 64)
	if err != nil {
		return err
	}
	_, err = p.DB.Exec(`DELETE FROM appointment WHERE id=$1`, apptId)
	if err != nil {
		return err
	}
	return nil
}
