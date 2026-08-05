package chat

type ChatMediator interface {
	Send(autor User, message string)
	Join(user User) error
	Leave(user User) error
}
