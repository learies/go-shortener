package app

import (
	"log/slog"
	"net/http"
)

// App represents the application.
type App struct{}

// New creates a new App instance.
func New() *App {
	return &App{}
}

// Run starts the HTTP server.
func (a *App) Run() error {
	addr := "localhost:8080"

	slog.Info("Listening server", "address", addr)

	return http.ListenAndServe(addr, nil)
}
