package main

type Empty struct{}

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
