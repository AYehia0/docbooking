package confirmation

import (
	"docbooking/pkg/event"
	"docbooking/pkg/logger"
)

type ConfirmationModule struct {
	bus *event.Bus
	log *logger.Logger

	// to send notifications
	notifier Notifier
}

func NewConfirmationModule(eventBus *event.Bus, log *logger.Logger) *ConfirmationModule {

	notifier := NewLogNotifier()

	return &ConfirmationModule{
		log:      log,
		bus:      eventBus,
		notifier: notifier,
	}
}

func (m *ConfirmationModule) RegisterEventListeners() {
	m.bus.Subscribe("booking.appointment.booked", func(e event.Event) {
		m.log.Info("Received booking.appointment.booked event in the confirmation module: sending notification to the doctor")
	})

	m.bus.Subscribe("booking.appointment.cancelled", func(e event.Event) {
		m.log.Info("Received booking.appointment.cancelled event in the confirmation module: sending notification to the patient")
	})
}

func (m *ConfirmationModule) SendNotification(e event.Event) {
	notification, ok := e.Payload.(Notification)
	if !ok {
		return
	}
	m.notifier.SendNotification(notification)
}
