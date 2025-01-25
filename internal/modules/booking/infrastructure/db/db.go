package db

import (
	"docbooking/internal/modules/booking/domain/entity"
	"docbooking/pkg/uuid"
	"fmt"
	"time"
)

type AppointmentRepo struct {
	data map[string]entity.Appointment
}

// usually not the best thing in the world to have 2 dbs for one module, but for the sake of simplicity
// and it will act as kinda of shared db, duplication is fine for now
type AvailabilityRepo struct {
	data map[string]entity.Slot
}

func NewAppointmentRepo() *AppointmentRepo {
	return &AppointmentRepo{
		data: make(map[string]entity.Appointment),
	}
}

func NewAvailabilityRepo() *AvailabilityRepo {
	return &AvailabilityRepo{
		data: make(map[string]entity.Slot),
	}
}

func (r *AvailabilityRepo) GetAvailableSlots(doctorID uuid.UUID) ([]entity.Slot, error) {
	var slots []entity.Slot
	for _, slot := range r.data {
		if slot.DoctorID == doctorID {
			slots = append(slots, slot)
		}
	}
	return slots, nil
}

func (r *AvailabilityRepo) AddSlot(doctorID uuid.UUID, start, end time.Time) error {
	slot := entity.Slot{
		ID:        uuid.New(),
		DoctorID:  doctorID,
		StartTime: start,
		EndTime:   end,
	}
	r.data[slot.ID.String()] = slot
	return nil
}

func (r *AppointmentRepo) Save(appointment entity.Appointment) error {
	if _, ok := r.data[appointment.ID.String()]; ok {
		return entity.ErrAppointmentAlreadyExists
	}

	r.data[appointment.ID.String()] = appointment
	fmt.Println("Saving an appointment!")
	return nil
}

func (r *AppointmentRepo) FindByID(id uuid.UUID) (entity.Appointment, error) {
	appointment, ok := r.data[id.String()]
	if ok {
		return appointment, nil
	}
	return entity.Appointment{}, entity.ErrAppointmentNotFound
}
