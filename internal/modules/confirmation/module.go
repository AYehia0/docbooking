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
	m.bus.Subscribe("confirmation.notification", m.SendNotification)
}

func (m *ConfirmationModule) SendNotification(e event.Event) {
	notification, ok := e.Payload.(Notification)
	if !ok {
		return
	}
	m.notifier.SendNotification(notification)
}
