package appointment

import (
	"docbooking/internal/modules/appointment/core/domain"
	appointmentservice "docbooking/internal/modules/appointment/core/service/appointment_service"
	"docbooking/internal/modules/appointment/handler"
	"docbooking/internal/modules/appointment/repo"
	"docbooking/pkg/event"
	"docbooking/pkg/logger"
	"net/http"
)

type AppointmentModule struct {
	Handler http.Handler
	bus     *event.Bus
}

func NewAppointmentModule(eventBus *event.Bus, logger *logger.Logger) *AppointmentModule {
	appointmentRepo := repo.NewAppointmentRepo()
	appointmentService := appointmentservice.NewAppointmentService(appointmentRepo)
	appointmentHandler := handler.NewAppointmentHandler(appointmentService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{doctor_id}/", appointmentHandler.GetDoctorAppointments)
	mux.HandleFunc("PUT /{appointment_id}/status/", appointmentHandler.UpdateAppointmentStatus)

	registerEvents(eventBus, logger, appointmentService)

	logger.Info("Appointment module is mounted on /appointments")
	logger.Info("GET /{doctor_id}")
	logger.Info("PUT /{appointment_id}/status")

	return &AppointmentModule{
		Handler: mux,
		bus:     eventBus,
	}
}

func registerEvents(eventBus *event.Bus, logger *logger.Logger, uc *appointmentservice.Service) {
	eventBus.Subscribe("booking.appointment.created", func(e event.Event) {
		appt := e.Payload.(event.CreateAppointmentEvent).Appointment
		logger.Infof("Received booking.appointment.created event in the appointment module, with payload: %+v", appt)

		uc.AddAppointment(domain.Appointment{
			ID:          appt.ID,
			SlotID:      appt.SlotID,
			PatientID:   appt.PatientID,
			DoctorID:    appt.DoctorID,
			PatientName: appt.PatientName,
			ReservedAt:  appt.ReservedAt,
			Status:      domain.AppointmentStatusPending,
		})
	})
}
