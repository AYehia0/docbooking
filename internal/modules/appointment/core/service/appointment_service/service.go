package appointmentservice

import (
	"docbooking/internal/modules/appointment/core/domain"
	"docbooking/internal/modules/appointment/core/port"

	"github.com/google/uuid"
)

type Service struct {
	repo port.AppointmentRepo
}

func NewAppointmentService(repo port.AppointmentRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetDoctorAppointments(doctorID uuid.UUID) ([]domain.Appointment, error) {
	return s.repo.GetDoctorAppointments(doctorID)
}

func (s *Service) UpdateAppointmentStatus(appointmentID uuid.UUID, status domain.AppointmentStatus) error {
	return s.repo.UpdateAppointmentStatus(appointmentID, status)
}

func (s *Service) AddAppointment(appointment domain.Appointment) error {
	return s.repo.AddAppointment(appointment)
}
