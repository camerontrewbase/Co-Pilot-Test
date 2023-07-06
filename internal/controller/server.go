package controller

import (
	"github.com/camerontrewbase/Co-Pilot-Test/internal/controller/handlers"
	"net/http"
)

// Create a factory function which returns a pointer to http server and an error and sets the address of the server to 8080. Use a Mux router to handle the routes. Have a route called /hello which returns a string "Hello World" and a route called /goodbye which returns a string "Goodbye World"
func NewServer() (*http.Server, error) {
	mux := http.NewServeMux()
	handlers.AddRoutes(mux)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}, nil
}
