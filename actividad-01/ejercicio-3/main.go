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
	fmt.Println("------ normal users scenario: ------")
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
	fmt.Println("------ user joins and leaves scenario: ------")
	room := rooms.NewGroupChatRoom()

	users := createUsers(4, room)

	for i := range 3 {
		user := users[i]
		room.Join(user)
	}

	// El usuario 3 no se ha unido a la sala, el mensaje no se envía
	users[3].SendMessage("Hello, everyone!")

	room.Join(users[3])

	users[3].SendMessage("I was talking on mute")
	users[1].SendMessage("Sorry, I gotta go!")

	room.Leave(users[1])

	// El usuario 1 no debería ver este mensaje
	users[2].SendMessage("Too bad :c")
}

func main() {
	fmt.Println()
	normalUsersScenario()
	fmt.Println()
	userJoinsAndLeavesScenario()
	fmt.Println()
}
