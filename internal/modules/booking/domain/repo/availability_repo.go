package repo

import (
	"docbooking/internal/modules/booking/domain/entity"
	"time"

	"github.com/google/uuid"
)

type AvailabilityRepo interface {
	GetAvailableSlots(doctorID uuid.UUID) ([]entity.Slot, error)
	AddSlot(doctorID uuid.UUID, start, end time.Time) error
}
