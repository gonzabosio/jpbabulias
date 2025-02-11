package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gonzabosio/jpbabulias/db/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func (h *Handler) AddAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	reqbody := new(model.Appointment)
	if err := json.NewDecoder(r.Body).Decode(&reqbody); err != nil {
		writeJSON(w, map[string]string{
			"message":    "Appointment data decodification error",
			"error_dets": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	if err := validate.Struct(reqbody); err != nil {
		errors := err.(validator.ValidationErrors)
		writeJSON(w, map[string]string{
			"message":    "Appointment data validation error",
			"error_dets": errors.Error(),
		}, http.StatusBadRequest)
		return
	}
	if err := h.rp.SaveAppointment(reqbody); err != nil {
		writeJSON(w, map[string]string{
			"message":    "Failed to save new appointment",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]interface{}{
		"message": "Appointment saved",
		"appt_id": reqbody.ID,
	}, http.StatusCreated)
}

func (h *Handler) GetAppointmentsByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	if userId == "" {
		writeJSON(w, map[string]string{
			"message":    "Not user id was provided",
			"error_dets": "user_id query is empty or invalid",
		}, http.StatusBadRequest)
		return
	}
	appts, err := h.rp.ReadAppointments(userId)
	if err != nil {
		writeJSON(w, map[string]string{
			"message":    fmt.Sprintf("Failed to read appointments"),
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]interface{}{
		"message":      "Appointments retrieved",
		"appointments": appts,
	}, http.StatusOK)
}

func (h *Handler) DeleteAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	apptIdStr := chi.URLParam(r, "appt_id")
	if apptIdStr == "" {
		writeJSON(w, map[string]string{
			"message":    "Not appointment id provided",
			"error_dets": "appt_id query is empty or invalid",
		}, http.StatusBadRequest)
		return
	}

	if err := h.rp.DeleteAppointment(apptIdStr); err != nil {
		writeJSON(w, map[string]string{
			"message":    fmt.Sprintf("Failed to delete appointment %s", apptIdStr),
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]interface{}{
		"message": "Appointment deleted",
	}, http.StatusOK)
}
