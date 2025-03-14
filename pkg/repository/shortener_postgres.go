package repository

import (
	"errors"
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
	query := fmt.Sprintf("UPDATE %s SET short_url = $1 WHERE id = $2", urlsTable)
	res, err := r.db.Exec(query, shortURL, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("nothing updated")
	}

	return nil
}

func (r *ShortenerPostgres) GetLongURL(id int) (string, error) {
	query := fmt.Sprintf("SELECT long_url FROM %s WHERE id = $1", urlsTable)
	var longURL string
	if err := r.db.Get(&longURL, query, id); err != nil {
		return "", err
	}

	return longURL, nil
}
