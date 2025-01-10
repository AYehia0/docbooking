package handler

import "github.com/google/uuid"

type BookingRequest struct {
	SlotID      uuid.UUID `json:"slot_id"`
	PatientID   uuid.UUID `json:"patient_id"`
	PatientName string    `json:"patient_name"`
}
