package platforms

import (
	"ejercicio-2/notifications"
	"fmt"
)

type mobileDecorator struct {
	notifier notifications.Notifier
}

func (md mobileDecorator) Send(message string) {
	message = fmt.Sprintf("[mobile] %s", message)
	md.notifier.Send(message)
}

func NewMobileDecorator(notifier notifications.Notifier) notifications.NotifierDecorator {
	return mobileDecorator{
		notifier: notifier,
	}
}
