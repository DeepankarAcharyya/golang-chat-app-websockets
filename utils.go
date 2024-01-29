package main

import "github.com/gorilla/websocket"

type Payload struct {
	Chatroom_id string `json:"ChatroomId"`
	Username    string `json:"User"`
	Message     string `json:"Message"`
}

type ChatRoomManager struct {
	Chatroom_id string
	Subscribers []*websocket.Connection
}

