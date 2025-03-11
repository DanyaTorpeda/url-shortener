package repository

import (
	shortener "url-shortener"

	"github.com/jmoiron/sqlx"
)

const (
	urlsTable = "urls"
)

type Shortener interface {
	CreateLongURL(input shortener.URL) (int, error)
	AddShortURL(id int, shortURL string) error
}

type Repository struct {
	Shortener
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Shortener: NewShortenerPostgres(db),
	}
}
