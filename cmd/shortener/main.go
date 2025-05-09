package main

import (
	"log/slog"

	"github.com/learies/go-shortener/internal/app"
	"github.com/learies/go-shortener/internal/config"
)

func main() {
	cfg := config.MustLoadConfig()

	app := app.NewApp(cfg)

	if err := app.Run(); err != nil {
		slog.Error("Could not start server", "error", err)
	}
}
