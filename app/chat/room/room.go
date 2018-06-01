package room

// Room distribute channel that holds incoming messages
type Room struct {
	forward chan []byte // updates other channels
	join    chan *Client
	leave   chan *Client
	clients map[*Client]Empty
}
