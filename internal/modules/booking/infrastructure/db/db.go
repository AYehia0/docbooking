package db

import (
	"docbooking/internal/modules/booking/domain/entity"
	"fmt"
	"github.com/google/uuid"
)

type AppointmentRepo struct {
	data map[string]entity.Appointment
}

func NewAppointmentRepo() *AppointmentRepo {
	return &AppointmentRepo{
		data: make(map[string]entity.Appointment),
	}
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
