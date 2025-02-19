package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func WriteJSON(w http.ResponseWriter, payload interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func hashPassword(password string) (hashedPassword string, err error) {
	bHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bHash), nil
}

func SetCookie(w http.ResponseWriter, name, value string, maxAge int) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		MaxAge:   maxAge,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})
}
