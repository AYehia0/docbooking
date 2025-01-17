package appointmentservice

import (
	"docbooking/internal/modules/appointment/core/domain"
	"docbooking/internal/modules/appointment/core/port"

	"github.com/google/uuid"
)

type service struct {
	repo port.AppointmentRepo
}

func NewAppointmentService(repo port.AppointmentRepo) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetDoctorAppointments(doctorID uuid.UUID) ([]domain.Appointment, error) {
	return s.repo.GetDoctorAppointments(doctorID)
}

func (s *service) UpdateAppointmentStatus(appointmentID uuid.UUID, status domain.AppointmentStatus) error {
	return s.repo.UpdateAppointmentStatus(appointmentID, status)
}
