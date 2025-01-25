package entity

import (
	"docbooking/pkg/uuid"
	"time"
)

// Appointment represents a booking appointment entity.
type Appointment struct {
	ID          uuid.UUID
	SlotID      uuid.UUID
	PatientID   uuid.UUID
	DoctorID    uuid.UUID
	StartTime   time.Time
	EndTime     time.Time
	PatientName string
	ReservedAt  time.Time
}
