package handler

// expose the handlers for the booking module

import (
	"docbooking/internal/modules/booking/application/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type BookingHandler struct {
	usecases *usecase.UseCases
}

func NewHTTPBookingHandler(usecases *usecase.UseCases) *BookingHandler {
	return &BookingHandler{
		usecases: usecases,
	}
}

func (h *BookingHandler) BookSlot(w http.ResponseWriter, r *http.Request) {

	doctorID, err := extractDoctorID(r.URL.Path)
	if err != nil {
		http.Error(w, "invalid doctor id", http.StatusBadRequest)
		return
	}

	var bookingRequest BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&bookingRequest); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if bookingRequest.SlotID == uuid.Nil || bookingRequest.PatientID == uuid.Nil || bookingRequest.PatientName == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	err = h.usecases.BookAppointmentUseCase.Execute(bookingRequest.SlotID, doctorID, bookingRequest.PatientID, bookingRequest.PatientName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *BookingHandler) GetDoctorAppointments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting doctor appointments: from patient's perspective")
	doctorID, err := extractDoctorID(r.URL.Path)
	if err != nil {
		http.Error(w, "invalid doctor id", http.StatusBadRequest)
		return
	}

	appointments, err := h.usecases.GetAppointmentsUseCase.Execute(doctorID)
	if err != nil {
		http.Error(w, "error getting appointments", http.StatusInternalServerError)
		return
	}

	var resp []AppointmentDto

	for _, appoint := range appointments {
		resp = append(resp, AppointmentDto{
			SlotID:    appoint.SlotID,
			StartTime: appoint.StartTime,
			EndTime:   appoint.EndTime,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func extractDoctorID(path string) (uuid.UUID, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 || parts[0] != "booking" {
		return uuid.Nil, http.ErrNotSupported
	}
	return uuid.Parse(parts[1])
}
