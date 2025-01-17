package availability

import (
	"docbooking/internal/modules/availability/handler"
	"docbooking/internal/modules/availability/repo"
	"docbooking/internal/modules/availability/service"
	"docbooking/pkg/event"
	"docbooking/pkg/logger"
	"net/http"

	"github.com/google/uuid"
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

	// Register the event handlers
	registerAvailabilityEvents(eventBus, availabilityService, logger)

	return &AvailabilityModule{
		Handler: mux,
		bus:     eventBus,
	}
}

// register the event handlers
func registerAvailabilityEvents(eventBus *event.Bus, service *service.AvailabilityService, logger *logger.Logger) {
	eventBus.Subscribe("availability.slot.created", func(e event.Event) {
		logger.Info("Received GetAvailableSlots event")
		doctorID, ok := e.Payload.(uuid.UUID)
		if !ok {
			logger.Error("Invalid payload for GetAvailableSlots event")
			return
		}

		slots, err := service.GetDoctorAvailabilitySlots(doctorID)
		if err != nil {
			logger.Errorf("Error fetching available slots: ", err)
			return
		}

		// Publish the response back
		eventBus.Publish(event.Event{
			Name:    "AvailableSlotsFetched",
			Payload: slots,
		})
	})

	// for testing
	eventBus.Subscribe("booking.appointment.created", func(e event.Event) {
		logger.Info("Received booking.appointment.created event")
	})
}
