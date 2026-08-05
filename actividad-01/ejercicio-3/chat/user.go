package chat

import "fmt"

type User struct {
	username string
	room     ChatMediator
}

func (u User) SendMessage(message string) {
	u.room.Send(u, message)
}

func (u User) ReceiveMessage(message string) {
	fmt.Printf("User [%s] received message \"%s\"\n", u.username, message)
}

func (u User) Username() string {
	return u.username
}

func NewUser(username string, room ChatMediator) User {
	return User{
		username: username,
		room:     room,
	}
}
