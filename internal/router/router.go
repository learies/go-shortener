package router

import (
	"net/http"
)

// Router represents an HTTP request router.
type Router interface {
	http.Handler
}

type serveMux struct {
	mux *http.ServeMux
}

// New creates a new Router instance.
func New() Router {
	mux := http.NewServeMux()

	return &serveMux{
		mux: mux,
	}
}

func (s *serveMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
