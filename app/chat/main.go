package main

import (
	"github.com/gorilla/websocket"
)

type Empty struct{}

// Room distribute channel that holds incoming messages
type Room struct {
	forward chan []byte // updates other channels
	join    chan *Client
	leave   chan *Client
	clients map[*Client]Empty
}

// bit hacky
func (r *Room) run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = Empty{}
		case client := <-r.leave:
			client.socket.Close()
			delete(r.clients, client)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				err := client.send(msg)
				if err != nil {
					r.leave <- client
				}
			}
		}
	}
}

// Client represents a user
type Client struct {
	socket *websocket.Conn
	room   *Room
}

func (c *Client) read() {
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			break
		}
		c.room.forward <- msg
	}
	c.socket.Close()
}

func (c *Client) send(msg []byte) error {
	return c.socket.WriteMessage(websocket.TextMessage, msg)
}
