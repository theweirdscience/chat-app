package handler

import "net/http"

// DefaultFileMW ...
type DefaultFileMW struct {
	handler http.Handler
}

func (mw DefaultFileMW) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.RequestURI == "/" {
		r.RequestURI = "/index.html"
	}

	mw.handler.ServeHTTP(w, r)

}

func main() {

	http.Handle("/", DefaultFileMW{
		handler: http.FileServer(http.Dir("./client")),
	})

	http.ListenAndServe(":8080", nil)

}
