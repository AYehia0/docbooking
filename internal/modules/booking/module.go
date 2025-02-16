package booking

import (
	"docbooking/internal/modules/booking/application/usecase"
	"docbooking/internal/modules/booking/domain/service"
	"docbooking/internal/modules/booking/infrastructure/db"
	"docbooking/internal/modules/booking/infrastructure/handler"
	"docbooking/pkg/event"
	"docbooking/pkg/logger"
	"net/http"
)

type BookModule struct {
	Handler http.Handler
	bus     *event.Bus
}

func NewBookModule(eventBus *event.Bus, logger *logger.Logger) *BookModule {
	bookingRepo := db.NewAppointmentRepo()
	availabilityRepo := db.NewAvailabilityRepo()

	bookingService := service.NewAppointmentService(bookingRepo, availabilityRepo, eventBus)
	availabilityService := service.NewAvailabilityService(availabilityRepo, eventBus)

	uc := &usecase.UseCases{
		GetAppointmentsUseCase: usecase.NewGetAppointmentsUseCase(bookingService),
		BookAppointmentUseCase: usecase.NewBookAppointmentUseCase(bookingService),
		AddSlotUseCase:         usecase.NewAddSlotUseCase(availabilityService),
	}

	bookingHandler := handler.NewHTTPBookingHandler(uc)

	registerEvents(eventBus, logger, uc)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /{doctor_id}/", bookingHandler.BookSlot)
	mux.HandleFunc("GET /{doctor_id}/", bookingHandler.GetDoctorAppointments)

	logger.Info("Booking module is mounted on /booking")
	logger.Info("POST /{doctor_id}/")

	return &BookModule{
		Handler: mux,
		bus:     eventBus,
	}
}

func registerEvents(eventBus *event.Bus, logger *logger.Logger, uc *usecase.UseCases) {
	eventBus.Subscribe("availability.slot.created", func(e event.Event) {
		slot := e.Payload.(event.CreateSlotEvent).Slot
		logger.Infof("Received availability.slot.created event: adding the slot as an appointment in the booking module: %v", slot)

		// TODO: handle errors
		uc.AddSlotUseCase.Execute(slot.DoctorID, slot.StartTime, slot.EndTime)
	})
}
