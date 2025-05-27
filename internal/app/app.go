package app

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/learies/go-shortener/internal/config"
	"github.com/learies/go-shortener/internal/handler"
	"github.com/learies/go-shortener/internal/router"
	"github.com/learies/go-shortener/internal/service"
	"github.com/learies/go-shortener/internal/storage/postgres"
)

type App struct {
	Config *config.Config
	Router *router.Router
}

func NewApp(cfg *config.Config) *App {
	shortener := service.NewURLShortener()

	storage, err := postgres.NewStorage(cfg)
	if err != nil {
		slog.Error("failed to create storage", "error", err)
		os.Exit(1)
	}

	h := handler.NewHandler(shortener, storage)
	r := router.NewRouter(h)

	return &App{
		Config: cfg,
		Router: r,
	}
}

func (a *App) Run() error {
	addr := a.Config.Server.Host + ":" + a.Config.Server.Port

	slog.Info("Starting server", "address", addr)
	return http.ListenAndServe(addr, a.Router.Mux)
}
