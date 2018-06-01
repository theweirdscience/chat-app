package handler

import "net/http"

// AppHandler handles the index, duh ;)
func main() {

	http.Handle("/", http.FileServer(http.Dir("./client/index.html")))
	http.ListenAndServe(":3000", nil)

}
