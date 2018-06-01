package room

// Client ...
type Client interface {
	Receive(chan<- []byte)
	Send([]byte) error
	Close() error
}

// Empty ...
type Empty struct{}

// Room distribute channel that holds incoming messages
type Room struct {
	forward chan []byte // updates other channels
	join    chan Client
	leave   chan Client
	clients map[Client]Empty
}

// Subscribe ...
func (r *Room) Subscribe(c Client) {
	r.join <- c
	c.Receive(r.forward)
}

// Run ...
func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = Empty{}
		case client := <-r.leave:
			delete(r.clients, client)
			client.Close()
		case msg := <-r.forward:
			for client := range r.clients {
				err := client.Send(msg)
				if err != nil {
					r.leave <- client
				}
			}
		}
	}
}
