package main

import (
	"github.com/gorilla/websocket"
)

// Room distribute channel that holds incoming messages
type Room struct {
	forward chan []byte // updates other channels
	join    chan *Client
	leave   chan *Client
	clients map[*Client]bool
}

// Client represents a user
type Client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *Room
}

func (c *Client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *Client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
