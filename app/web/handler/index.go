package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/theweirdscience/chat-app/lib/message"
)

// DefaultFileMW ...
type DefaultFileMW struct {
	handler http.Handler
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan message.Message)
var upgrader = websocket.Upgrader{}

func (mw DefaultFileMW) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.RequestURI == "/" {
		r.RequestURI = "/index.html"
	}

	mw.handler.ServeHTTP(w, r)

}

// HandleIndex ...
func HandleIndex() {

	go handleMessages()

	http.Handle("/", DefaultFileMW{
		handler: http.FileServer(http.Dir("./client")),
	})

	http.HandleFunc("/ws", handleConnections)

	http.ListenAndServe(":8080", nil)

}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg message.Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}

}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
