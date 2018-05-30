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

// bit hacky
func (r *Room) run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send the message
				default:
					// failed to send
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
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
