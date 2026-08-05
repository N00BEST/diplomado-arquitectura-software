package notifications

import "ejercicio-2/platforms"

type baseNotifier struct {
	platform platforms.Displayer
}

func (bs *baseNotifier) SetPlatform(platform platforms.Displayer) {
	bs.platform = platform
}
