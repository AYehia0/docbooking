package domain

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
	PatientName string
	ReservedAt  time.Time
	Status      AppointmentStatus
}

// either cancelled, completed, or pending
type AppointmentStatus string

const (
	AppointmentStatusCancelled AppointmentStatus = "cancelled"
	AppointmentStatusCompleted AppointmentStatus = "completed"
	AppointmentStatusPending   AppointmentStatus = "pending"
)
