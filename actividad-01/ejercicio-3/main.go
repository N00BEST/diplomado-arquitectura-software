package main

import (
	"ejercicio-3/chat"
	"ejercicio-3/rooms"
	"fmt"
)

func createUsers(count int, room chat.ChatMediator) []chat.User {
	users := make([]chat.User, 0, count)

	for i := range count {
		username := fmt.Sprintf("User %d", i)
		users = append(users, chat.NewUser(username, room))
	}

	return users
}

func normalUsersScenario() {
	room := rooms.NewGroupChatRoom()

	users := createUsers(3, room)

	for _, user := range users {
		room.Join(user)
	}

	users[1].SendMessage("Hello, world!")
	users[2].SendMessage("It's alive!!")
	users[0].SendMessage("Glad it works")
}

func userJoinsAndLeavesScenario() {
	room := rooms.NewGroupChatRoom()

	users := createUsers(5, room)

	for _, user := range users {
		room.Join(user)
	}

	users[3].SendMessage("Hello, everyone!")
	users[1].SendMessage("Sorry, I gotta go!")

	room.Leave(users[1])

	users[4].SendMessage("Too bad :c")
}

func main() {
	normalUsersScenario()
	userJoinsAndLeavesScenario()
}
