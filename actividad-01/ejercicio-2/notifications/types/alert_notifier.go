package types

import (
	"ejercicio-2/notifications"
	"fmt"
)

type alertNotifier struct {
}

func (mn alertNotifier) Send(message string) {
	fmt.Println("[alert]", message)
}

func NewAlertNotifier() notifications.Notifier {
	return alertNotifier{}
}
