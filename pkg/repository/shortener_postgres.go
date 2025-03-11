package repository

import (
	"fmt"
	shortener "url-shortener"

	"github.com/jmoiron/sqlx"
)

type ShortenerPostgres struct {
	db *sqlx.DB
}

func NewShortenerPostgres(db *sqlx.DB) *ShortenerPostgres {
	return &ShortenerPostgres{db: db}
}

func (r *ShortenerPostgres) CreateLongURL(input shortener.URL) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (long_url) values($1) RETURNING id", urlsTable)
	row := r.db.QueryRow(query, input.LongURL)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ShortenerPostgres) AddShortURL(id int, shortURL string) error {

}
