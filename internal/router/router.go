package router

import (
	"net/http"

	"github.com/learies/go-shortener/internal/handler"
)

// Router represents an HTTP request router.
type Router interface {
	http.Handler
}

type serveMux struct {
	mux *http.ServeMux
}

// New creates a new Router instance.
func New(shortenerHandler *handler.Handler) Router {
	mux := http.NewServeMux()

	mux.Handle("POST /", http.HandlerFunc(shortenerHandler.CreateShortURL))

	return &serveMux{
		mux: mux,
	}
}

func (s *serveMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
