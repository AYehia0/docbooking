package repo

import (
	"docbooking/internal/modules/appointment/core/domain"
	"docbooking/pkg/uuid"
)

type appointmentRepo struct {
	appointments map[string][]domain.Appointment
}

func NewAppointmentRepo() *appointmentRepo {
	return &appointmentRepo{
		appointments: make(map[string][]domain.Appointment),
	}
}

func (r *appointmentRepo) GetDoctorAppointments(doctorID uuid.UUID) ([]domain.Appointment, error) {
	appointments, ok := r.appointments[doctorID.String()]
	if !ok {
		return nil, nil
	}
	return appointments, nil
}

func (r *appointmentRepo) UpdateAppointmentStatus(appointmentID uuid.UUID, status domain.AppointmentStatus) error {
	for _, appointments := range r.appointments {
		for i := range appointments {
			if appointments[i].ID == appointmentID {
				appointments[i].Status = status
				return nil
			}
		}
	}
	return nil
}

func (r *appointmentRepo) AddAppointment(appointment domain.Appointment) error {
	r.appointments[appointment.DoctorID.String()] = append(r.appointments[appointment.DoctorID.String()], appointment)
	return nil
}
