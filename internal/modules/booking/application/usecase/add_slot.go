package usecase

import (
	"docbooking/internal/modules/booking/domain/service"
	"docbooking/pkg/uuid"
	"time"
)

type AddSlotUseCase struct {
	availabilityService *service.AvailabilityService
}

func NewAddSlotUseCase(availabilityService *service.AvailabilityService) *AddSlotUseCase {
	return &AddSlotUseCase{
		availabilityService: availabilityService,
	}
}

func (uc *AddSlotUseCase) Execute(doctorID uuid.UUID, start, end time.Time) error {
	return uc.availabilityService.AddSlot(doctorID, start, end)
}
