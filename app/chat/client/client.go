package client

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