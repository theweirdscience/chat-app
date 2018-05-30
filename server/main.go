package main

import (
	"log"
	"net/http"
)

func main() {

	handleHome()

}

func handleHome() {

	http.Handle("/", http.FileServer(http.Dir("./client")))

	if err := http.ListenAndServe(":8080", nil); err != nil {

		log.Fatal("ListenAndServe:", err)

	}

}
