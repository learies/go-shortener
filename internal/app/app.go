package app

import (
	"log/slog"
	"net/http"

	"github.com/learies/go-shortener/internal/config"
	"github.com/learies/go-shortener/internal/handler"
	"github.com/learies/go-shortener/internal/router"
	"github.com/learies/go-shortener/internal/service"
)

type App struct {
	Config *config.Config
	Router *router.Router
}

func NewApp(cfg *config.Config) *App {
	shortener := service.NewURLShortener()

	h := handler.NewHandler(shortener)
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
