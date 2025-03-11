package shortener

import "time"

type URL struct {
	ID        int
	LongURL   string
	ShortURL  string
	CreatedAt time.Time
}
