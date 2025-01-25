package service

import (
	"docbooking/internal/modules/booking/domain/entity"
	"docbooking/internal/modules/booking/domain/repo"
	"docbooking/pkg/event"
	"docbooking/pkg/uuid"
	"errors"
	"time"
)

type AppointmentService struct {
	availabilityRepo repo.AvailabilityRepo
	appointmentRepo  repo.AppointmentRepo
	bus              *event.Bus
}

func NewAppointmentService(appointmentRepo repo.AppointmentRepo, availabilityRepo repo.AvailabilityRepo, bus *event.Bus) *AppointmentService {
	return &AppointmentService{
		appointmentRepo:  appointmentRepo,
		availabilityRepo: availabilityRepo,
		bus:              bus,
	}
}

func (s *AppointmentService) BookAppointment(slotID, doctorID, patientID uuid.UUID, patientName string) error {
	_, err := s.appointmentRepo.FindByID(slotID)
	if err == nil {
		return errors.New("slot is already booked")
	}

	appointment := entity.Appointment{
		ID:          uuid.NewRandom(),
		SlotID:      slotID,
		PatientID:   patientID,
		DoctorID:    doctorID,
		PatientName: patientName,
		ReservedAt:  time.Now(),
	}

	if err := s.appointmentRepo.Save(appointment); err != nil {
		return err
	}

	s.bus.Publish(event.Event{
		Name: "booking.appointment.created",
		Payload: event.CreateAppointmentEvent{
			Appointment: event.Appointment(appointment),
		},
	})

	return nil
}

func (s *AppointmentService) GetDoctorAppointments(doctorID uuid.UUID) ([]entity.Appointment, error) {
	var docAppointments []entity.Appointment

	slots, err := s.availabilityRepo.GetAvailableSlots(doctorID)
	if err != nil {
		return nil, err
	}

	for _, slot := range slots {
		if slot.IsReserved || isSlotExpired(slot) {
			continue
		}
		docAppointments = append(docAppointments, entity.Appointment{
			SlotID:    slot.ID,
			StartTime: slot.StartTime,
			EndTime:   slot.EndTime,
		})
	}

	return docAppointments, nil
}

func isSlotExpired(slot entity.Slot) bool {
	return slot.EndTime.Before(time.Now())
}
