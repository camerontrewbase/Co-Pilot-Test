package internal

import (
	"github.com/camerontrewbase/Co-Pilot-Test/internal/controller"
	"net/http"
)

// Create a new App struct. Which contains a pointer to a http server
type App struct {
	Server *http.Server
}

// Create a function called run which serves the http server on port 8080
func (app *App) Run() error {
	return app.Server.ListenAndServe()
}

// Create a factory function to create and initialise a new App struct
func NewApp() (*App, error) {
	server, err := controller.NewServer()
	if err != nil {
		return nil, err
	}

	return &App{
		Server: server,
	}, nil
}
