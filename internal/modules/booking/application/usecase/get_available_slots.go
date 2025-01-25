package usecase

import (
	"docbooking/internal/modules/booking/domain/entity"
	"docbooking/internal/modules/booking/domain/service"
	"docbooking/pkg/uuid"
)

type GetAppointmentsUseCase struct {
	appointmentService *service.AppointmentService
}

func NewGetAppointmentsUseCase(appointmentService *service.AppointmentService) *GetAppointmentsUseCase {
	return &GetAppointmentsUseCase{
		appointmentService: appointmentService,
	}
}

func (uc *GetAppointmentsUseCase) Execute(doctorID uuid.UUID) ([]entity.Appointment, error) {
	return uc.appointmentService.GetDoctorAppointments(doctorID)
}
