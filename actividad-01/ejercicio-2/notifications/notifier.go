package notifications

type Notifier interface {
	Send(message string)
}
