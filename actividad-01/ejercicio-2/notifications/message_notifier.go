package notifications

import (
	"ejercicio-2/platforms"
)

type messageNotifier struct {
	baseNotifier
}

func (mn *messageNotifier) Send(message string) {
	mn.platform.Display("message", message)
}

func NewMessageNotifier(platform platforms.Displayer) Notifier {
	return &messageNotifier{
		baseNotifier: baseNotifier{
			platform: platform,
		},
	}
}
