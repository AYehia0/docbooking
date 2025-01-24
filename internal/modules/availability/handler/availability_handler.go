package handler

import (
	"docbooking/internal/modules/availability/model"
	"docbooking/internal/modules/availability/service"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type HTTPAvailabilityHandler struct {
	availabilityService *service.AvailabilityService
}

func NewHTTPAvailabilityHandler(service *service.AvailabilityService) *HTTPAvailabilityHandler {
	return &HTTPAvailabilityHandler{
		availabilityService: service,
	}
}

// GetDoctorAvailabilitySlots handles GET /availabilities/{doctorID}
func (h *HTTPAvailabilityHandler) GetDoctorAvailabilitySlots(w http.ResponseWriter, r *http.Request) {
	// Extract doctorID from URL
	doctorID, err := extractDoctorID(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Fetch slots
	slots, err := h.availabilityService.GetDoctorAvailabilitySlots(doctorID)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slots)
}

// AddDoctorAvailabilitySlots handles POST /availabilities/{doctorID}
func (h *HTTPAvailabilityHandler) AddDoctorAvailabilitySlots(w http.ResponseWriter, r *http.Request) {
	// Extract doctorID from URL
	doctorID, err := extractDoctorID(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse request body
	var slot model.Slot
	if err := json.NewDecoder(r.Body).Decode(&slot); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validate slot times
	if slot.StartTime.After(slot.EndTime) {
		http.Error(w, "start time cannot be after end time", http.StatusBadRequest)
		return
	}

	if slot.StartTime.Before(time.Now()) {
		http.Error(w, "start time cannot be in the past", http.StatusBadRequest)
		return
	}

	// Add slot
	err = h.availabilityService.AddDoctorAvailabilitySlots(doctorID, slot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "slot added successfully"})
}

func extractDoctorID(path string) (uuid.UUID, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 || parts[0] != "availabilities" {
		return uuid.Nil, http.ErrNotSupported
	}
	return uuid.Parse(parts[1])
}
