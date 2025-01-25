package port

import (
	"docbooking/internal/modules/appointment/core/domain"
	"docbooking/pkg/uuid"
)

// The DB port
type AppointmentRepo interface {
	GetDoctorAppointments(doctorID uuid.UUID) ([]domain.Appointment, error)
	UpdateAppointmentStatus(appointmentID uuid.UUID, status domain.AppointmentStatus) error
	AddAppointment(appointment domain.Appointment) error
}

// The service port
type AppointmentService interface {
	GetDoctorAppointments(doctorID uuid.UUID) ([]domain.Appointment, error)
	UpdateAppointmentStatus(appointmentID uuid.UUID, status domain.AppointmentStatus) error
	AddAppointment(appointment domain.Appointment) error
}
