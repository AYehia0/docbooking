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
	bookingService := service.NewAppointmentService(bookingRepo, eventBus)

	uc := &usecase.UseCases{
		GetAppointmentsUseCase: usecase.NewGetAppointmentsUseCase(bookingService),
		BookAppointmentUseCase: usecase.NewBookAppointmentUseCase(bookingService),
	}

	bookingHandler := handler.NewHTTPBookingHandler(uc)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /{doctor_id}/", bookingHandler.BookSlot)

	return &BookModule{
		Handler: mux,
		bus:     eventBus,
	}
}
