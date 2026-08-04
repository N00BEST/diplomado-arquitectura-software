package platforms

import (
	"ejercicio-2/notifications"
	"fmt"
)

type webDecorator struct {
	notifier notifications.Notifier
}

func (wd webDecorator) Send(message string) {
	message = fmt.Sprintf("[web] %s", message)
	wd.notifier.Send(message)
}

func NewWebDecorator(notifier notifications.Notifier) notifications.NotifierDecorator {
	return webDecorator{
		notifier: notifier,
	}
}
