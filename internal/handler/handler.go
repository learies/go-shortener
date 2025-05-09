package handler

import (
	"github.com/learies/go-shortener/internal/service"
)

type Handler struct {
	shortener service.Shortener
}

func NewHandler(shortener service.Shortener) *Handler {
	return &Handler{
		shortener: shortener,
	}
}
