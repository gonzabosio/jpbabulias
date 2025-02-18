package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
