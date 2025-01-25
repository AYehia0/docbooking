package repo

import (
	"docbooking/internal/modules/booking/domain/entity"
	"docbooking/pkg/uuid"
)

type AppointmentRepo interface {
	Save(appointment entity.Appointment) error
	FindByID(id uuid.UUID) (entity.Appointment, error)
}
