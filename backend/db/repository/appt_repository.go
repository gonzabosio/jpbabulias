package repository

import (
	"strconv"
	"time"

	"github.com/gonzabosio/jpbabulias/db/model"
)

type AppointmentRepository interface {
	SaveAppointment(appt *model.Appointment) error
	ReadAppointmentsByUserId(userIdStr string) ([]model.AppointmentList, error)
	ReadAppointmentsByDay(day string) ([]model.AppointmentList, error)
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

func (p *PostgreService) ReadAppointmentsByUserId(userIdStr string) ([]model.AppointmentList, error) {
	var appts []model.AppointmentList
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
	return appts, nil
}

func (p *PostgreService) ReadAppointmentsByDay(day string) ([]model.AppointmentList, error) {
	parsedDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		return nil, err
	}
	start := parsedDate.Format("2006-01-02 15:04:05")
	end := parsedDate.Add(24 * time.Hour).Format("2006-01-02 15:04:05")
	rows, err := p.DB.Query(`SELECT id, appt_date, subject FROM appointment WHERE appt_date >= $1 and appt_date <= $2`, start, end)
	if err != nil {
		return nil, err
	}
	var apptList []model.AppointmentList
	for rows.Next() {
		var appt model.AppointmentList
		rows.Scan(&appt.ID, &appt.ApptDate, &appt.Subject)
		apptList = append(apptList, appt)
	}
	return apptList, nil
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
