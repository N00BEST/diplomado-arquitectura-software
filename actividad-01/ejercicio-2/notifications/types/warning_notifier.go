package types

import (
	"ejercicio-2/notifications"
	"fmt"
)

type warningNotifier struct {
}

func (mn warningNotifier) Send(message string) {
	fmt.Println("[warning]", message)
}

func NewWarningNotifier() notifications.Notifier {
	return warningNotifier{}
}
