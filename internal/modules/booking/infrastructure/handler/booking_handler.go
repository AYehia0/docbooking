package handler

// expose the handlers for the booking module

import (
	"docbooking/internal/modules/booking/application/usecase"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
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
	fmt.Println("Booking a slot")
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

	err := h.usecases.BookAppointmentUseCase.Execute(bookingRequest.SlotID, bookingRequest.PatientID, bookingRequest.PatientName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
