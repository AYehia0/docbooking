package service

import (
	"docbooking/internal/modules/booking/domain/entity"
	"docbooking/internal/modules/booking/domain/repo"
	"docbooking/pkg/event"
	"errors"
	"time"

	"github.com/google/uuid"
)

type AppointmentService struct {
	repo repo.AppointmentRepo
	bus  *event.Bus
}

func NewAppointmentService(repo repo.AppointmentRepo, bus *event.Bus) *AppointmentService {
	return &AppointmentService{
		repo: repo,
		bus:  bus,
	}
}

func (s *AppointmentService) BookAppointment(slotID, patientID uuid.UUID, patientName string) error {
	// check if the slot is free to book and not already booked
	_, err := s.repo.FindByID(slotID)
	if err == nil {
		return errors.New("slot is already booked")
	}

	appointmentId, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	appointment := entity.Appointment{
		ID:          appointmentId,
		SlotID:      slotID,
		PatientID:   patientID,
		PatientName: patientName,
		ReservedAt:  time.Now(),
	}

	if err := s.repo.Save(appointment); err != nil {
		return err
	}

	// Publish the booking event
	s.bus.Publish(event.Event{
		Name:    "booking.appointment.created",
		Payload: appointment,
	})

	return nil
}

func (s *AppointmentService) GetDoctorAppointments(doctorID uuid.UUID) ([]entity.Appointment, error) {
	// use the event bus to get the appointments
	var docAppointments []entity.Appointment
	s.bus.Subscribe("availability.slot.created", func(e event.Event) {
		appointments := e.Payload.([]entity.Appointment)

		// filter the appointments by doctorID
		// currently all the appointments are related to one doctor
		docAppointments = append(docAppointments, appointments...)
	})

	return docAppointments, nil
}
