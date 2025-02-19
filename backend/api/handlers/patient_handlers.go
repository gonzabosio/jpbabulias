package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gonzabosio/jpbabulias/db/model"
)

func (h *Handler) GetPatientsByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to parse user ID",
			"error_dets": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	patientsList, err := h.rp.ReadPatientsByUserId(userId)
	if err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to get user patients list",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	WriteJSON(w, map[string]interface{}{
		"message":  "Patients retrieved",
		"patients": patientsList,
	}, http.StatusOK)
}

func (h *Handler) AddPatientHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody model.Patient
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Patient data decodification error",
			"error_dets": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&reqBody); err != nil {
		errors := err.(validator.ValidationErrors)
		WriteJSON(w, map[string]string{
			"message":    "Patient data validation error",
			"error_dets": errors.Error(),
		}, http.StatusBadRequest)
		return
	}
	if err := h.rp.SavePatient(&reqBody); err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to save new patient",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	WriteJSON(w, map[string]string{
		"message":    "Patient created",
		"patient_id": reqBody.ID,
	}, http.StatusCreated)
}
