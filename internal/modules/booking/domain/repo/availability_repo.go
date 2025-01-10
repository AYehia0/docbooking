package repo

import (
	"docbooking/internal/modules/booking/domain/entity"

	"github.com/google/uuid"
)

type AvailabilityRepo interface {
	GetAvailableSlots(doctorID uuid.UUID) ([]entity.Slot, error)
}
