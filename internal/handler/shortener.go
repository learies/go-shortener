package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/learies/go-shortener/internal/dto"
)

func (h *Handler) CreateShortLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "can't read body", http.StatusInternalServerError)
			return
		}

		originalURL := string(body)

		shortURL, err := h.shortener.GenerateShortURL(originalURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := dto.ShortenURLResponse{
			OriginalURL: originalURL,
			ShortURL:    shortURL,
		}

		if err := h.storage.Add(response); err != nil {
			http.Error(w, "failed to save URL", http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "can't encode JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}
}
