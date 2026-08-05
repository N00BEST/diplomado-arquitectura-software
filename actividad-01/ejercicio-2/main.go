package main

import (
	"ejercicio-2/notifications"
	"ejercicio-2/notifications/platforms"
	"ejercicio-2/notifications/types"
)

func notify(notification notifications.Notifier, message string) {
	notification.Send(message)
}

func vanillaScenario() {
	alert := types.NewAlertNotifier()

	notify(alert, "sending vanilla notification")
}

func singlePlatformScenario() {
	confirmation := types.NewConfirmationNotifier()
	web := platforms.NewWebDecorator(confirmation)

	notify(web, "sending notification to web")
}

func multiPlatformScenario() {
	warning := types.NewWarningNotifier()
	web := platforms.NewWebDecorator(warning)
	mobile := platforms.NewMobileDecorator(web)
	desktop := platforms.NewDesktopDecorator(mobile)

	notify(desktop, "new sign-in detected in your account")
}

func main() {
	vanillaScenario()
	singlePlatformScenario()
	multiPlatformScenario()
}
