package entity

import (
	"time"

	"github.com/google/uuid"
)

// The Slot entity represents a time slot for a doctor's availability, should be the same as the one in availability module,
// it's being defined here again just to follow the clean architecture depedency direction
type Slot struct {
	ID         uuid.UUID `json:"id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	IsReserved bool      `json:"is_reserved"`
	DoctorID   uuid.UUID `json:"doctor_id"`
	DoctorName string    `json:"doctor_name"`
	Cost       float64   `json:"cost"`
}
