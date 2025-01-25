package handler

import (
	"docbooking/internal/modules/appointment/core/domain"
	"docbooking/internal/modules/appointment/core/port"
	"docbooking/pkg/uuid"
	"encoding/json"
	"fmt"
	"net/http"
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
	fmt.Println("GetDoctorAppointments")
	doctorID, err := uuid.Parse(r.URL.Query().Get("doctor_id"))
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
	fmt.Println("UpdateAppointmentStatus")
	var req struct {
		AppointmentID uuid.UUID `json:"appointment_id"`
		Status        string    `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	if err := h.service.UpdateAppointmentStatus(req.AppointmentID, st); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
