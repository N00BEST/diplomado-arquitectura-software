package notifications

import (
	"ejercicio-2/platforms"
)

type alertNotifier struct {
	baseNotifier
}

func (an *alertNotifier) Send(message string) {
	an.platform.Display("alert", message)
}

func NewAlertNotifier(platform platforms.Displayer) Notifier {
	return &alertNotifier{
		baseNotifier: baseNotifier{
			platform: platform,
		},
	}
}
