package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

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
	domain := os.Getenv("DOMAIN")
	isLocal := false
	if domain == "localhost" {
		isLocal = true
	}
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: false,
		MaxAge:   maxAge,
		Path:     "/",
		Domain:   domain,
		Expires:  time.Now().Add(time.Duration(maxAge) * time.Second),
		SameSite: func() http.SameSite {
			if isLocal {
				return http.SameSiteLaxMode
			}
			return http.SameSiteNoneMode
		}(),
		Secure: !isLocal,
	})
}
