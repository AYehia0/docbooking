package handler

import (
	"docbooking/internal/modules/appointment/core/domain"
	"docbooking/internal/modules/appointment/core/port"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type AppointmentHandler struct {
	service port.AppointmentService
}

func NewAppointmentHandler(service port.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		service: service,
	}
}

func (h *AppointmentHandler) GetDoctorAppointments(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[2]
	doctorID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "invalid doctor_id", http.StatusBadRequest)
		return
	}

	appointments, err := h.service.GetDoctorAppointments(doctorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(appointments) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	json.NewEncoder(w).Encode(appointments)
}

func (h *AppointmentHandler) UpdateAppointmentStatus(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	appointmentID, err := extractAppointmentID(r.URL.Path)
	if err != nil {
		http.Error(w, "invalid appointment_id", http.StatusBadRequest)
		return
	}

	// ensure the status is valid
	switch req.Status {
	case "cancelled", "completed", "pending":
	default:
		http.Error(w, "invalid status", http.StatusBadRequest)
		return
	}
	// convert req.Status to port.AppointmentStatus
	st := domain.AppointmentStatus(req.Status)

	if err := h.service.UpdateAppointmentStatus(appointmentID, st); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func extractAppointmentID(path string) (uuid.UUID, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 || parts[0] != "appointments" {
		return uuid.Nil, http.ErrNotSupported
	}
	return uuid.Parse(parts[1])
}
