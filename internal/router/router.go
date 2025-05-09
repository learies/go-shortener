package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/learies/go-shortener/internal/handler"
)

type Router struct {
	*chi.Mux
	Handler *handler.Handler
}

func NewRouter(h *handler.Handler) *Router {
	r := &Router{
		Mux:     chi.NewRouter(),
		Handler: h,
	}

	r.setupMiddleware()
	r.setupRoutes()

	return r
}

func (r *Router) setupMiddleware() {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))
}

func (r *Router) setupRoutes() {
	r.Route("/api/v1", func(router chi.Router) {
		router.Post("/", r.Handler.CreateShortLink())
	})
}
