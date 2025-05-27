package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

type Shortener interface {
	GenerateShortURL(url string) (string, error)
}

type URLShortener struct{}

func NewURLShortener() *URLShortener {
	return &URLShortener{}
}

func (us *URLShortener) GenerateShortURL(url string) (string, error) {
	if url == "" {
		return "", errors.New("empty URL")
	}

	hasher := sha256.New()
	hasher.Write([]byte(url))
	hash := hasher.Sum(nil)

	encoded := base64.URLEncoding.EncodeToString(hash)
	if len(encoded) < 8 {
		return "", errors.New("generated hash is too short")
	}

	return encoded[:8], nil
}
