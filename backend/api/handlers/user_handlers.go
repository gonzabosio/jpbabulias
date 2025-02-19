package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gonzabosio/jpbabulias/api/token"
	"github.com/gonzabosio/jpbabulias/db/model"
)

func (h *Handler) UserSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody model.UserSignUp
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		WriteJSON(w, map[string]string{
			"message":    "User signup data decodification error",
			"error_dets": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&reqBody); err != nil {
		errors := err.(validator.ValidationErrors)
		WriteJSON(w, map[string]string{
			"message":    "User signup data validation error",
			"error_dets": errors.Error(),
		}, http.StatusBadRequest)
		return
	}
	hashpw, err := hashPassword(reqBody.Password)
	reqBody.Password = hashpw
	err = h.rp.SaveUser(&reqBody)
	if err != nil {
		WriteJSON(w, map[string]string{
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

	// credentials
	accTkn, err := token.GenerateAccessToken(userData.ID)
	if err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to generate access token",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rfTkn, err := token.GenerateRefreshToken(userData.ID)
	if err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to generate refresh token",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	SetCookie(w, "access_token", accTkn, token.ATMaxAgeStd)
	SetCookie(w, "refresh_token", rfTkn, token.RTMaxAgeStd)

	WriteJSON(w, map[string]interface{}{
		"message":   "User registered",
		"user_data": userData,
	}, http.StatusCreated)
}

func (h *Handler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody model.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		WriteJSON(w, map[string]string{
			"message":    "User signup data decodification error",
			"error_dets": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&reqBody); err != nil {
		errors := err.(validator.ValidationErrors)
		WriteJSON(w, map[string]string{
			"message":    "User signup data validation error",
			"error_dets": errors.Error(),
		}, http.StatusBadRequest)
		return
	}
	userData, err := h.rp.VerifyUser(&reqBody)
	if err != nil {
		if strings.Contains(err.Error(), "comparation") {
			WriteJSON(w, map[string]string{
				"message":    "Contraseña invalida o incorrecta",
				"error_dets": err.Error(),
			}, http.StatusBadRequest)
			return
		} else if err.Error() == "no email match" {
			WriteJSON(w, map[string]string{
				"message":    "El correo no existe",
				"error_dets": err.Error(),
			}, http.StatusBadRequest)
			return
		} else {
			WriteJSON(w, map[string]string{
				"message":    "Inicio de sesión fallido. Revise los campos e inténtelo de nuevos",
				"error_dets": err.Error(),
			}, http.StatusInternalServerError)
			return
		}
	}

	// credentials
	accTkn, err := token.GenerateAccessToken(userData.ID)
	if err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to generate access token",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rfTkn, err := token.GenerateRefreshToken(userData.ID)
	if err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to generate refresh token",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	SetCookie(w, "access_token", accTkn, token.ATMaxAgeStd)
	SetCookie(w, "refresh_token", rfTkn, token.RTMaxAgeStd)

	WriteJSON(w, map[string]interface{}{
		"message":   "User logged in",
		"user_data": userData,
	}, http.StatusOK)
}

func (h *Handler) RefreshAccessTokenHandler(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	SetCookie(w, "access_token", "", -1)
	SetCookie(w, "refresh_token", "", -1)
	WriteJSON(w, map[string]string{
		"message": "User logged out",
	}, http.StatusOK)
}
