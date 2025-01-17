package appointment

import (
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

	// log the endpoints
	logger.Info("Appointment module is mounted on /appointments")
	logger.Info("GET /{doctor_id}")
	logger.Info("PUT /{appointment_id}/status")

	return &AppointmentModule{
		Handler: mux,
		bus:     eventBus,
	}
}
