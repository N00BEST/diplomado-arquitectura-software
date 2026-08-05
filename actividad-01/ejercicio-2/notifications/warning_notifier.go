package notifications

import (
	"ejercicio-2/platforms"
)

type warningNotifier struct {
	baseNotifier
}

func (wn *warningNotifier) Send(message string) {
	wn.platform.Display("warning", message)
}

func NewWarningNotifier(platform platforms.Displayer) Notifier {
	return &warningNotifier{
		baseNotifier: baseNotifier{
			platform: platform,
		},
	}
}
