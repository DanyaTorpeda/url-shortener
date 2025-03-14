package service

import (
	shortener "url-shortener"
	"url-shortener/pkg/repository"
)

type Shortener interface {
	CreateShortURL(input shortener.URL) (string, error)
	GetLongURL(shortURL string) (string, error)
}

type Service struct {
	Shortener
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Shortener: NewShortenerService(repo.Shortener),
	}
}
