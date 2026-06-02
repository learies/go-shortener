package handler

import "net/http"

// Handler handles HTTP requests.
type Handler struct{}

// New creates a new Handler instance.
func New() *Handler {
	return &Handler{}
}

// CreateShortURL handles URL shortening requests.
// Example:
//
//	curl -X POST http://localhost:8080/ \
//		 -H "Content-Type: text/plain" \
//		 --data-binary "https://practicum.yandex.ru/"
func (h *Handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Hello, World"))
}
