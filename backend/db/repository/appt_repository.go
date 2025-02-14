package repository

import (
	"strconv"
	"time"

	"github.com/gonzabosio/jpbabulias/db/model"
)

type AppointmentRepository interface {
	SaveAppointment(appt *model.Appointment) error
	ReadAppointmentsByUserId(userIdStr string) (*[]model.AppointmentList, error)
	ReadAppointmentsByDay(day string) (*[]string, error)
	ReadFullyBookedDates() (*[]string, error)
	DeleteAppointment(userIdStrr string) error
}

var _ AppointmentRepository = (*PostgreService)(nil)

func (p *PostgreService) SaveAppointment(appt *model.Appointment) error {
	var apptID int
	userIdNum, err := strconv.ParseInt(appt.UserID, 10, 64)
	if err != nil {
		return err
	}
	err = p.DB.QueryRow(`INSERT INTO appointment(appt_date, subject, user_id) VALUES($1, $2, $3) RETURNING id`, appt.ApptDate, appt.Subject, userIdNum).Scan(&apptID)
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
	rows, err := p.DB.Query(`SELECT id, appt_date, subject FROM appointment WHERE user_id=$1`, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var appt model.AppointmentList
		if err := rows.Scan(&appt.ID, &appt.ApptDate, &appt.Subject); err != nil {
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
	for rows.Next() {
		var appt time.Time
		rows.Scan(&appt)
		apptList = append(apptList, appt)
	}
	var hoursList []string
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
		rows.Scan(&date)
		fullyBookedDates = append(fullyBookedDates, date)
	}
	return &fullyBookedDates, nil
}

func (p *PostgreService) DeleteAppointment(userIdStr string) error {
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return err
	}
	_, err = p.DB.Exec(`DELETE FROM appointment WHERE id=$1`, userId)
	if err != nil {
		return err
	}
	return nil
}
