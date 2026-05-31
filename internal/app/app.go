package app

import (
	"log/slog"
	"net/http"

	"github.com/learies/go-shortener/internal/router"
)

// App represents the application.
type App struct {
	router router.Router
}

// New creates a new App instance.
func New() *App {
	return &App{
		router: router.New(),
	}
}

// Run starts the HTTP server.
func (a *App) Run() error {
	addr := "localhost:8080"

	slog.Info("Listening server", "address", addr)

	return http.ListenAndServe(addr, a.router)
}
