package client

import (
	"github.com/gorilla/websocket"
)

// Client represents a user
type Client struct {
	socket *websocket.Conn
}

// Receive ...
func (c *Client) Receive(messages chan<- []byte) {

	go func() {
		for {
			_, msg, err := c.socket.ReadMessage()

			if err != nil {
				break
			}

			messages <- msg
		}
		c.socket.Close()

	}()

}

// Send ...
func (c *Client) Send(msg []byte) error {

	return c.socket.WriteMessage(websocket.TextMessage, msg)

}

// Close ...
func (c *Client) Close() {

	return c.socket.Close()

}
