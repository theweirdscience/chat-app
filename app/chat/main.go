package main

import (
	"github.com/theweirdscience/room"
)

func (r *room.Room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = struct{}
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
