package handlers

import "net/http"

// Create a function which takes a new http server mux and adds routes to it
func AddRoutes(mux *http.ServeMux) {
	// Create an endpoint of /hello which will take a GET parameter of name and return a string "Hello <name>"
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		w.Write([]byte("Hello " + name))
	})

	mux.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Goodbye World"))
	})
}
