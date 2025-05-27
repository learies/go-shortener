package handler

import (
	"github.com/learies/go-shortener/internal/service"
	"github.com/learies/go-shortener/internal/storage"
)

type Handler struct {
	shortener service.Shortener
	storage   storage.Storage
}

func NewHandler(shortener service.Shortener, storage storage.Storage) *Handler {
	return &Handler{
		shortener: shortener,
		storage:   storage,
	}
}
