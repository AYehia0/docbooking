package service

import (
	"docbooking/internal/modules/booking/domain/entity"
	"docbooking/internal/modules/booking/domain/repo"
	"docbooking/pkg/event"
	"time"

	"github.com/google/uuid"
)

type AvailabilityService struct {
	repo repo.AvailabilityRepo
	bus  *event.Bus
}

func NewAvailabilityService(repo repo.AvailabilityRepo, bus *event.Bus) *AvailabilityService {
	return &AvailabilityService{
		repo: repo,
		bus:  bus,
	}
}

func (s *AvailabilityService) GetAvailableSlots(doctorID uuid.UUID) ([]entity.Slot, error) {
	return s.repo.GetAvailableSlots(doctorID)
}

func (s *AvailabilityService) AddSlot(doctorID uuid.UUID, start, end time.Time) error {
	return s.repo.AddSlot(doctorID, start, end)
}
