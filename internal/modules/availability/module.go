package availability

import (
	"docbooking/internal/modules/availability/handler"
	"docbooking/internal/modules/availability/repo"
	"docbooking/internal/modules/availability/service"
	"docbooking/pkg/event"
	"docbooking/pkg/logger"
	"net/http"
)

type AvailabilityModule struct {
	Handler http.Handler
	bus     *event.Bus
}

func NewAvailabilityModule(eventBus *event.Bus, logger *logger.Logger) *AvailabilityModule {
	availabilityRepo := repo.NewAvailabilityRepo()
	availabilityService := service.NewAvailabilityService(eventBus, availabilityRepo)
	availabilityHandler := handler.NewHTTPAvailabilityHandler(availabilityService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{doctor_id}/", availabilityHandler.GetDoctorAvailabilitySlots)
	mux.HandleFunc("POST /{doctor_id}/", availabilityHandler.AddDoctorAvailabilitySlots)

	// log the endpoints
	logger.Info("Availability module is mounted on /availabilities")
	logger.Info("GET /{doctor_id}/")
	logger.Info("POST /{doctor_id}/")

	return &AvailabilityModule{
		Handler: mux,
		bus:     eventBus,
	}
}
