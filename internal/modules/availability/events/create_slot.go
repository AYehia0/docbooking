package events

import (
	"docbooking/internal/modules/availability/model"

	"github.com/google/uuid"
)

type CreateSlotEvent struct {
	DoctorID uuid.UUID  `json:"doctor_id"`
	Slot     model.Slot `json:"slot"`
}
