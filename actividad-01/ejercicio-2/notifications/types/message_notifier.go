package types

import (
	"ejercicio-2/notifications"
	"fmt"
)

type messageNotifier struct {
}

func (mn messageNotifier) Send(message string) {
	fmt.Println("[message]", message)
}

func NewMessageNotifier() notifications.Notifier {
	return messageNotifier{}
}
