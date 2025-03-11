package service

import (
	shortener "url-shortener"
	"url-shortener/pkg/repository"
)

type Shortener interface {
	Create(input shortener.URL) (int, error)
}

type Service struct {
	Shortener
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
