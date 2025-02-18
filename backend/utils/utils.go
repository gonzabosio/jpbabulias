package utils

import "golang.org/x/crypto/bcrypt"

func ComparePassword(pwHash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(pwHash), []byte(password)); err != nil {
		return err
	}
	return nil
}
