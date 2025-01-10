package repo

// in memeory repo for the availability

import (
	"docbooking/internal/modules/availability/model"
	"errors"

	"github.com/google/uuid"
)

type AvailabilityRepository interface {
	GetDoctorAvailabilitySlots(doctorID uuid.UUID) ([]model.Slot, error)
	AddDoctorAvailabilitySlot(doctorID uuid.UUID, slots model.Slot) error
}

type availabilityRepo struct {
	availabilitySlots map[uuid.UUID][]model.Slot
}

func NewAvailabilityRepo() AvailabilityRepository {
	return &availabilityRepo{
		availabilitySlots: make(map[uuid.UUID][]model.Slot),
	}
}

func (r *availabilityRepo) GetDoctorAvailabilitySlots(doctorID uuid.UUID) ([]model.Slot, error) {
	slots, ok := r.availabilitySlots[doctorID]
	if !ok {
		return nil, errors.New("no slots found")
	}

	return slots, nil
}

func (r *availabilityRepo) AddDoctorAvailabilitySlot(doctorID uuid.UUID, slot model.Slot) error {
	if slot.StartTime.After(slot.EndTime) {
		return errors.New("start time should be before end time")
	}
	r.availabilitySlots[doctorID] = append(r.availabilitySlots[doctorID], slot)
	return nil
}
