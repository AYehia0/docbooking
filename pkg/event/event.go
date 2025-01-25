package event

import (
	"docbooking/pkg/uuid"
	"time"
)

type Event struct {
	Name    string
	Payload interface{}
}

type CreateSlotEvent struct {
	DoctorID uuid.UUID
	Slot     Slot
}

type CreateAppointmentEvent struct {
	Appointment Appointment
}

// Shared domain types: don't know where to put them yet
type Slot struct {
	ID         uuid.UUID
	StartTime  time.Time
	EndTime    time.Time
	IsReserved bool
	DoctorID   uuid.UUID
	DoctorName string
	Cost       float64
}

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
