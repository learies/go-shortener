package storage

import (
	"github.com/learies/go-shortener/internal/dto"
)

type Storage interface {
	Add(dto.ShortenURLResponse) error
}
