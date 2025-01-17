package confirmation

import "fmt"

type Payload any

type Notification struct {
	Payload Payload
	Target  string
}

type Notifier interface {
	SendNotification(notification Notification)
}

type LogNotifier struct {
}

func (n *LogNotifier) SendNotification(notification Notification) {
	fmt.Printf("Sending notification using log notifier to %s with payload %v\n", notification.Target, notification.Payload)
}

func NewLogNotifier() *LogNotifier {
	return &LogNotifier{}
}
