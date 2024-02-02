package main

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type MessageUnit struct {
	Username string
	Message  string
}

type ChatRoomManager struct {
	Chatroom_id string
	Subscribers map[string]*websocket.Conn
	ModLock     sync.Mutex
}

func NewChatRoom() *ChatRoomManager {
	return &ChatRoomManager{
		Subscribers: make(map[string]*websocket.Conn),
	}
}

func (ChatRoom *ChatRoomManager) addNewParticipant(User string, NewParticipantConn *websocket.Conn) (bool, error) {
	ChatRoom.ModLock.Lock()
	defer ChatRoom.ModLock.Unlock()

	if _, ok := ChatRoom.Subscribers[User]; ok {
		// Check if a user with the username already exists
		message := fmt.Sprintf("User with the following username (%s)  already exists.", User)
		return false, errors.New(message)
	}

	ChatRoom.Subscribers[User] = NewParticipantConn
	return true, nil
}

func (ChatRoom *ChatRoomManager) removeParticipant(User string) (bool, error) {
	ChatRoom.ModLock.Lock()
	defer ChatRoom.ModLock.Unlock()

	if _, ok := ChatRoom.Subscribers[User]; ok {
		delete(ChatRoom.Subscribers, User)
		return true, nil
	}

	return false, errors.New("User doesn't exist.")
}

func (chatRoom *ChatRoomManager) SendMessage(msg MessageUnit) {
	for _, sub := range chatRoom.Subscribers {
		// code to send the messages
	}
}
