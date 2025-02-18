package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gonzabosio/jpbabulias/db/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) AddAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	reqbody := new(model.Appointment)
	if err := json.NewDecoder(r.Body).Decode(&reqbody); err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Appointment data decodification error",
			"error_dets": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	if err := validate.Struct(reqbody); err != nil {
		errors := err.(validator.ValidationErrors)
		WriteJSON(w, map[string]string{
			"message":    "Appointment data validation error",
			"error_dets": errors.Error(),
		}, http.StatusBadRequest)
		return
	}
	if err := h.rp.SaveAppointment(reqbody); err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to save new appointment",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	WriteJSON(w, map[string]interface{}{
		"message": "Appointment saved",
		"appt_id": reqbody.ID,
	}, http.StatusCreated)
}

func (h *Handler) GetAppointmentsByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	if userId == "" {
		WriteJSON(w, map[string]string{
			"message":    "Not user id was provided",
			"error_dets": "user_id query is empty or invalid",
		}, http.StatusBadRequest)
		return
	}
	appts, err := h.rp.ReadAppointmentsByUserId(userId)
	if err != nil {
		WriteJSON(w, map[string]string{
			"message":    fmt.Sprintf("Failed to read appointments by user id"),
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	WriteJSON(w, map[string]interface{}{
		"message":      "Appointments retrieved",
		"appointments": &appts,
	}, http.StatusOK)
}

func (h *Handler) GetAppointmentsByDayHandler(w http.ResponseWriter, r *http.Request) {
	day := r.URL.Query().Get("date")
	if day == "" {
		WriteJSON(w, map[string]string{
			"message":    "Failed to get date",
			"error_dets": "empty query field",
		}, http.StatusBadRequest)
		return
	}

	appts, err := h.rp.ReadAppointmentsByDay(day)
	if err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to read appointments by day",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if len(*appts) == 0 {
		WriteJSON(w, map[string]interface{}{
			"message":      "Appointments retrieved",
			"appointments": []string{},
		}, http.StatusOK)
		return
	}
	WriteJSON(w, map[string]interface{}{
		"message":      "Appointments retrieved",
		"appointments": &appts,
	}, http.StatusOK)
}

func (h *Handler) GetFullyBookedDatesHandler(w http.ResponseWriter, r *http.Request) {
	dateList, err := h.rp.ReadFullyBookedDates()
	if err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to get the fully booked dates",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	WriteJSON(w, map[string]interface{}{
		"message":            "Fully booked dates retrieved",
		"fully_booked_dates": &dateList,
	}, http.StatusOK)
}

func (h *Handler) DeleteAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	apptIdStr := chi.URLParam(r, "appt_id")
	if apptIdStr == "" {
		WriteJSON(w, map[string]string{
			"message":    "Not appointment id provided",
			"error_dets": "appt_id query is empty or invalid",
		}, http.StatusBadRequest)
		return
	}

	if err := h.rp.DeleteAppointment(apptIdStr); err != nil {
		WriteJSON(w, map[string]string{
			"message":    fmt.Sprintf("Failed to delete appointment %s", apptIdStr),
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	WriteJSON(w, map[string]interface{}{
		"message": "Appointment deleted",
	}, http.StatusOK)
}
