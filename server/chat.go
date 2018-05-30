package main

import (
	"github.com/gorilla/websocket"
)

// Room distribute channel that holds incoming messages
type Room struct {
	distribute chan []byte
}

// Client represents a user
type Client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *Room
}
