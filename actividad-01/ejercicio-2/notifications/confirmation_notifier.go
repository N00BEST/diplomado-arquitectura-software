package notifications

import (
	"ejercicio-2/platforms"
)

type confirmationNotifier struct {
	baseNotifier
}

func (cn *confirmationNotifier) Send(message string) {
	cn.platform.Display("confirmation", message)
}

func NewConfirmationNotifier(platform platforms.Displayer) Notifier {
	return &confirmationNotifier{
		baseNotifier: baseNotifier{
			platform: platform,
		},
	}
}
