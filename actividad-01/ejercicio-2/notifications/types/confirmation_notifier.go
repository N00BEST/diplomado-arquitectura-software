package types

import (
	"ejercicio-2/notifications"
	"fmt"
)

type confirmationNotifier struct {
}

func (mn confirmationNotifier) Send(message string) {
	fmt.Println("[confirmation]", message)
}

func NewConfirmationNotifier() notifications.Notifier {
	return confirmationNotifier{}
}
