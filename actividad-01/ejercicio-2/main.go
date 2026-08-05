package main

import (
	"ejercicio-2/notifications"
	"ejercicio-2/platforms"
	"fmt"
)

func singlePlatformScenario() {
	fmt.Println("------ single platform scenario: ------")
	web := platforms.NewWebDisplayer()
	confirmation := notifications.NewConfirmationNotifier(web)

	confirmation.Send("sending notification to web")
}

func multiNotificationScenario() {
	fmt.Println("------ multi notification scenario: ------")
	mobile := platforms.NewMobileDisplayer()
	warning := notifications.NewWarningNotifier(mobile)
	message := notifications.NewMessageNotifier(mobile)

	warning.Send("poor network connection detected")
	message.Send("device is back online")
}

func multiPlatformScenario() {
	fmt.Println("------ multi platform scenario: ------")
	all := []platforms.Displayer{
		platforms.NewDesktopDisplayer(),
		platforms.NewMobileDisplayer(),
		platforms.NewWebDisplayer(),
	}

	alert := notifications.NewAlertNotifier(nil)

	for _, platform := range all {
		alert.SetPlatform(platform)
		alert.Send("new sign-in detected into your account")
	}
}

func main() {
	fmt.Println("\n")
	singlePlatformScenario()
	fmt.Println("\n")
	multiNotificationScenario()
	fmt.Println("\n")
	multiPlatformScenario()
	fmt.Println("\n")
}
