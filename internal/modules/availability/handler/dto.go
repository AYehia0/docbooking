package handler

import (
	"docbooking/pkg/uuid"
	"time"
)

// DTO for adding a slot
type AddSlotRequest struct {
	DoctorID  uuid.UUID `json:"doctor_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Cost      float64   `json:"cost"`
}

// DTO for slot response
type SlotResponse struct {
	ID          uuid.UUID `json:"id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	IsAvailable bool      `json:"is_available"`
	DoctorID    uuid.UUID `json:"doctor_id"`
	DoctorName  string    `json:"doctor_name"`
	Cost        float64   `json:"cost"`
}
