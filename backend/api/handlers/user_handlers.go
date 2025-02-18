package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gonzabosio/jpbabulias/db/model"
)

func (h *Handler) UserSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody model.UserSignUp
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		writeJSON(w, map[string]string{
			"message":    "User signup data decodification error",
			"error_dets": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&reqBody); err != nil {
		errors := err.(validator.ValidationErrors)
		writeJSON(w, map[string]string{
			"message":    "User signup data validation error",
			"error_dets": errors.Error(),
		}, http.StatusBadRequest)
		return
	}
	hashpw, err := hashPassword(reqBody.Password)
	reqBody.Password = hashpw
	err = h.rp.SaveUser(&reqBody)
	if err != nil {
		writeJSON(w, map[string]string{
			"message":    "Failed to save user/patient",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	userData := &model.UserData{
		ID:    reqBody.ID,
		Email: reqBody.Email,
		Admin: reqBody.Admin,
	}
	writeJSON(w, map[string]interface{}{
		"message":   "User registered",
		"user_data": userData,
	}, http.StatusCreated)
}

func (h *Handler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody model.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		writeJSON(w, map[string]string{
			"message":    "User signup data decodification error",
			"error_dets": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&reqBody); err != nil {
		errors := err.(validator.ValidationErrors)
		writeJSON(w, map[string]string{
			"message":    "User signup data validation error",
			"error_dets": errors.Error(),
		}, http.StatusBadRequest)
		return
	}
	userData, err := h.rp.VerifyUser(&reqBody)
	if err != nil {
		if strings.Contains(err.Error(), "comparation") {
			writeJSON(w, map[string]string{
				"message":    "Contraseña invalida o incorrecta",
				"error_dets": err.Error(),
			}, http.StatusBadRequest)
			return
		} else if err.Error() == "no email match" {
			writeJSON(w, map[string]string{
				"message":    "El correo no existe",
				"error_dets": err.Error(),
			}, http.StatusBadRequest)
			return
		} else {
			writeJSON(w, map[string]string{
				"message":    "Inicio de sesión fallido. Revise los campos e inténtelo de nuevos",
				"error_dets": err.Error(),
			}, http.StatusInternalServerError)
			return
		}
	}
	writeJSON(w, map[string]interface{}{
		"message":   "User logged in",
		"user_data": userData,
	}, http.StatusOK)
}
