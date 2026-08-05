package notifications

import "ejercicio-2/platforms"

type Notifier interface {
	Send(message string)
	SetPlatform(platform platforms.Displayer)
}
