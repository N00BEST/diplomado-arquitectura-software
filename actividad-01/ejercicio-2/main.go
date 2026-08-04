package main

import (
	"ejercicio-2/notifications/platforms"
	"ejercicio-2/notifications/types"
)

func vanillaScenario() {
	alert := types.NewAlertNotifier()

	alert.Send("sending vanilla notification")
}

func singlePlatformScenario() {
	confirmation := types.NewConfirmationNotifier()
	web := platforms.NewWebDecorator(confirmation)

	web.Send("sending notification to web")
}

func multiPlatformScenario() {
	warning := types.NewWarningNotifier()
	web := platforms.NewWebDecorator(warning)
	mobile := platforms.NewMobileDecorator(web)
	desktop := platforms.NewDesktopDecorator(mobile)

	desktop.Send("new sign-in detected in your account")
}

func main() {
	vanillaScenario()
	singlePlatformScenario()
	multiPlatformScenario()
}
