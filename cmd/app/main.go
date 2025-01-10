package main

import (
	"docbooking/internal/modules/availability"
	"docbooking/internal/modules/booking"
	"docbooking/pkg/event"
	"docbooking/pkg/logger"
	"net/http"
)

func main() {
	// the modular monolith entry point used to initialize all the modules
	eventBus := event.NewEventBus()
	log := logger.NewLogger()

	// initialize the availability module
	availabilityModule := availability.NewAvailabilityModule(eventBus, log)
	bookingModule := booking.NewBookModule(eventBus, log)

	// initializing the server
	mux := http.NewServeMux()
	mux.Handle("/availabilities/", availabilityModule.Handler)
	mux.Handle("/booking/", bookingModule.Handler)

	// running the server
	log.Infof("Server is running on port %s", "8080")
	http.ListenAndServe(":8080", mux)
}
