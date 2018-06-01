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

	// defer to regular handler
	mw.handler.ServeHTTP(w, r)

}

// AppHandler handles the index, duh ;)
func main() {

	http.Handle("/", DefaultFileMW{
		handler: http.FileServer(http.Dir("./client/dist")),
	})

	http.ListenAndServe(":8080", nil)

}
