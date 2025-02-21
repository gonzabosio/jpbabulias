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
		"patients": &patientsList,
	}, http.StatusOK)
}

func (h *Handler) AddPatientHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody model.InsertPatient
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

func (h *Handler) EditPatientDataHandler(w http.ResponseWriter, r *http.Request) {
	reqBody := new(model.UpdatePatient)
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Patient data decodification error",
			"error_dets": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	// if err := validate.Struct(&reqBody); err != nil {
	// 	errors := err.(validator.ValidationErrors)
	// 	WriteJSON(w, map[string]string{
	// 		"message":    "Patient data validation error",
	// 		"error_dets": errors.Error(),
	// 	}, http.StatusBadRequest)
	// 	return
	// }
	err := h.rp.EditPatientData(reqBody)
	if err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to edit patient data",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	WriteJSON(w, map[string]string{
		"message": "Patient data updated",
	}, http.StatusOK)
}

func (h *Handler) DeletePatientByIdHandler(w http.ResponseWriter, r *http.Request) {
	patientIdStr := chi.URLParam(r, "patient_id")
	patientId, err := strconv.Atoi(patientIdStr)
	if err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to delete patient",
			"error_dets": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	err = h.rp.DeletePatientById(patientId)
	if err != nil {
		if err.Error() == "patient have existing appointments" {
			WriteJSON(w, map[string]string{
				"message":    "Failed to delete patient",
				"error_dets": err.Error(),
			}, http.StatusConflict)
			return
		}
		WriteJSON(w, map[string]string{
			"message":    "Failed to delete patient",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	WriteJSON(w, map[string]string{
		"message": "Patient deleted",
	}, http.StatusOK)
}
