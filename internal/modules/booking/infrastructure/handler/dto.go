package handler

import (
	"time"

	"github.com/google/uuid"
)

type BookingRequest struct {
	SlotID      uuid.UUID `json:"slot_id"`
	PatientID   uuid.UUID `json:"patient_id"`
	PatientName string    `json:"patient_name"`
}

type AppointmentDto struct {
	SlotID    uuid.UUID `json:"slot_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
