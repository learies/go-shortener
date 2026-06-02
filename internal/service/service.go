package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"sync"
)

var ErrNotFound = errors.New("short URL not found")

// Shortener stores shortened URLs.
type Shortener struct {
	mu   sync.RWMutex
	urls map[string]string
}

// New creates a new Shortener instance.
func New() *Shortener {
	return &Shortener{
		urls: make(map[string]string),
	}
}

// Create saves the original URL and returns its short ID.
func (s *Shortener) Create(originalURL string) string {
	shortID := generateShortID(originalURL)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.urls[shortID] = originalURL

	return shortID
}

// Get returns the original URL by short ID.
func (s *Shortener) Get(id string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	originalURL, exists := s.urls[id]
	if !exists {
		return "", ErrNotFound
	}

	return originalURL, nil
}

func generateShortID(s string) string {
	hash := sha256.Sum256([]byte(s))

	encoded := base64.RawURLEncoding.EncodeToString(hash[:])

	return encoded[:8]
}
