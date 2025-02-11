package repository

import (
	"strconv"

	"github.com/gonzabosio/jpbabulias/db/model"
)

type AppointmentRepository interface {
	SaveAppointment(appt *model.Appointment) error
	ReadAppointments(userIdStr string) ([]model.AppointmentList, error)
	DeleteAppointment(userIdStrr string) error
}

var _ AppointmentRepository = (*PostgreService)(nil)

func (p *PostgreService) SaveAppointment(appt *model.Appointment) error {
	var apptID int
	userIdNum, err := strconv.ParseInt(appt.UserID, 10, 64)
	if err != nil {
		return err
	}
	err = p.DB.QueryRow(`INSERT INTO appointment(appt_date, user_id) VALUES($1, $2) RETURNING id`, appt.ApptDate, userIdNum).Scan(&apptID)
	if err != nil {
		return err
	}
	appt.ID = strconv.Itoa(apptID)
	return nil
}

func (p *PostgreService) ReadAppointments(userIdStr string) ([]model.AppointmentList, error) {
	var appts []model.AppointmentList
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return nil, err
	}
	rows, err := p.DB.Query(`SELECT id, appt_date FROM appointment WHERE user_id=$1`, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var appt model.AppointmentList
		if err := rows.Scan(&appt.ID, &appt.ApptDate); err != nil {
			return nil, err
		}
		appts = append(appts, appt)
	}
	return appts, nil
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
