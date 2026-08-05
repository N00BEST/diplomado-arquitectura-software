package rooms

import (
	"ejercicio-3/chat"
	"fmt"
)

type groupChatRoom struct {
	users map[string]chat.User
}

func (gcr *groupChatRoom) Send(author chat.User, message string) {
	if len(gcr.users) == 0 {
		return
	}

	currentUsername := author.Username()

	// El usuario no está en la sala, no puede enviar mensajes
	if _, exists := gcr.users[currentUsername]; !exists {
		return
	}

	for username, user := range gcr.users {
		if username == currentUsername {
			continue
		}

		msg := fmt.Sprintf("[%s]: %s", currentUsername, message)

		user.ReceiveMessage(msg)
	}
}

func (gcr *groupChatRoom) Join(user chat.User) error {
	username := user.Username()

	if _, exists := gcr.users[username]; exists {
		return fmt.Errorf("user is already in the group chat room")
	}

	gcr.users[username] = user
	return nil
}

func (gcr *groupChatRoom) Leave(user chat.User) error {
	username := user.Username()

	if _, exists := gcr.users[username]; !exists {
		return fmt.Errorf("user is not inside the group chat room")
	}

	delete(gcr.users, username)
	return nil
}

func NewGroupChatRoom() chat.ChatMediator {
	return &groupChatRoom{
		users: make(map[string]chat.User),
	}
}
