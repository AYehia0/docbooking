package service

import (
	"docbooking/internal/modules/availability/events"
	"docbooking/internal/modules/availability/model"
	"docbooking/internal/modules/availability/repo"
	"docbooking/pkg/event"

	"github.com/google/uuid"
)

type AvailabilityService struct {
	availabilityRepo repo.AvailabilityRepository

	// the event bus
	bus *event.Bus
}

func NewAvailabilityService(bus *event.Bus, availabilityRepo repo.AvailabilityRepository) *AvailabilityService {
	return &AvailabilityService{
		availabilityRepo: availabilityRepo,
		bus:              bus,
	}
}

// The doctor can view his availability slots
func (s *AvailabilityService) GetDoctorAvailabilitySlots(doctorID uuid.UUID) ([]model.Slot, error) {
	return s.availabilityRepo.GetDoctorAvailabilitySlots(doctorID)
}

// The doctor can add his availability slots
func (s *AvailabilityService) AddDoctorAvailabilitySlots(doctorID uuid.UUID, slot model.Slot) error {
	err := s.availabilityRepo.AddDoctorAvailabilitySlot(doctorID, slot)

	if err == nil {
		s.bus.Publish(event.Event{
			Name: "availability.slot.created",
			Payload: events.CreateSlotEvent{
				DoctorID: doctorID,
				Slot:     slot,
			},
		})
	}

	return err
}
