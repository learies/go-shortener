package handler

import (
	"errors"
	"io"
	"net/http"

	"github.com/learies/go-shortener/internal/service"
)

// ShortenerService describes URL shortening operations used by Handler.
type ShortenerService interface {
	Create(originalURL string) string
	Get(shortID string) (string, error)
}

// Handler handles HTTP requests.
type Handler struct {
	service ShortenerService
}

// New creates a new Handler instance.
func New(service ShortenerService) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateShortURL handles URL shortening requests.
// Example:
//
//	curl -X POST http://localhost:8080/ \
//		 -H "Content-Type: text/plain" \
//		 --data-binary "https://practicum.yandex.ru/"
func (h *Handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusRequestEntityTooLarge)
		return
	}

	originalURL := string(body)
	if originalURL == "" {
		http.Error(w, "empty URL", http.StatusBadRequest)
		return
	}

	shortID := h.service.Create(originalURL)

	shortURL := "http://localhost:8080/" + shortID

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))
}

// Redirect handles short URL redirect requests.
// Example:
//
//	curl -v http://localhost:8080/QrPnX5IU
func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("shortID")
	if id == "" {
		http.Error(w, "empty short URL id", http.StatusBadRequest)
		return
	}

	originalURL, err := h.service.Get(id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, "short URL not found", http.StatusNotFound)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
