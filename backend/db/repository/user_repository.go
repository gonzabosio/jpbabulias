package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gonzabosio/jpbabulias/db/model"
	"github.com/gonzabosio/jpbabulias/utils"
)

type UserRepository interface {
	SaveUser(user *model.UserSignUp) error
	VerifyUser(user *model.UserLogin) (*model.UserData, error)
}

var _ UserRepository = (*PostgreService)(nil)

func (p *PostgreService) SaveUser(user *model.UserSignUp) error {

	var userId int
	tx, err := p.DB.Begin()
	if err != nil {
		return fmt.Errorf("transaction instance creation error: %v", err)
	}
	row := tx.QueryRow(`INSERT INTO "user" (email, password, birth_date, admin) 
	VALUES($1, $2, $3, $4) RETURNING id;`, user.Email, user.Password, user.BirthDate, false)
	if err := row.Scan(&userId); err != nil {
		tx.Rollback()
		return fmt.Errorf("user creation error: %v", err)
	}

	userIdStr := strconv.Itoa(userId)
	user.ID = userIdStr
	_, err = tx.Exec(
		`INSERT INTO patient(first_name, last_name, phone_number, dni, health_insurance, main, user_id) 
		VALUES($1, $2, $3, $4, $5, $6, $7);`,
		user.FirstName, user.LastName, user.PhoneNumber, user.Dni, user.HealthInsurance, true, user.ID,
	)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("patient creation error: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}

func (p *PostgreService) VerifyUser(userReq *model.UserLogin) (*model.UserData, error) {
	var dbUser model.User
	row := p.DB.QueryRow(`SELECT * FROM "user" WHERE email=$1`, userReq.Email)
	if err := row.Scan(&dbUser.ID, &dbUser.Email, &dbUser.Password, &dbUser.BirthDate, &dbUser.Admin); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no email match")
		}
		return nil, err
	}
	if err := utils.ComparePassword(dbUser.Password, userReq.Password); err != nil {
		return nil, fmt.Errorf("password comparation failed: %v", err)
	}
	userData := &model.UserData{
		ID:    dbUser.ID,
		Email: dbUser.Email,
		Admin: dbUser.Admin,
	}
	return userData, nil
}
