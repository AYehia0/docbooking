package usecase

import (
	"docbooking/internal/modules/booking/domain/service"

	"github.com/google/uuid"
)

type BookAppointmentUseCase struct {
	appointmentService *service.AppointmentService
}

func NewBookAppointmentUseCase(appointmentService *service.AppointmentService) *BookAppointmentUseCase {
	return &BookAppointmentUseCase{
		appointmentService: appointmentService,
	}
}

func (uc *BookAppointmentUseCase) Execute(slotID, doctorID, patientID uuid.UUID, patientName string) error {
	return uc.appointmentService.BookAppointment(slotID, doctorID, patientID, patientName)
}
