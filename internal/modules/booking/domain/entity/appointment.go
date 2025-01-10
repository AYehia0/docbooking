package entity

import (
	"time"

	"github.com/google/uuid"
)

// Appointment represents a booking appointment entity.
type Appointment struct {
	ID          uuid.UUID
	SlotID      uuid.UUID
	PatientID   uuid.UUID
	PatientName string
	ReservedAt  time.Time
}
