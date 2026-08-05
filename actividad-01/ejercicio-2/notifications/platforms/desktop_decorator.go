package platforms

import (
	"ejercicio-2/notifications"
	"fmt"
)

type desktopDecorator struct {
	notifier notifications.Notifier
}

func (dd desktopDecorator) Send(message string) {
	message = fmt.Sprintf("[desktop] %s", message)
	dd.notifier.Send(message)
}

func NewDesktopDecorator(notifier notifications.Notifier) notifications.NotifierDecorator {
	return desktopDecorator{
		notifier: notifier,
	}
}
