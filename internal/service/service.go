package service

import (
	"sync"
)

// Shortener stores shortened URLs.
type Shortener struct {
	mu   sync.Mutex
	urls map[string]string
}

// New creates a new Shortener instance.
func New() *Shortener {
	return &Shortener{
		urls: make(map[string]string),
	}
}
