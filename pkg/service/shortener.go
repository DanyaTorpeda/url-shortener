package service

import (
	"fmt"
	"strconv"
	shortener "url-shortener"
	"url-shortener/pkg/repository"

	"github.com/jxskiss/base62"
	"github.com/spf13/viper"
)

type ShortenerService struct {
	repo repository.Shortener
}

func NewShortenerService(repo repository.Shortener) *ShortenerService {
	return &ShortenerService{repo: repo}
}

func (s *ShortenerService) Create(input shortener.URL) (int, error) {
	id, err := s.repo.CreateLongURL(input)
	if err != nil {
		return 0, err
	}

	convertedId := base62.Encode([]byte(strconv.Itoa(id)))

	shortURL := fmt.Sprintf("http://localhost:%s/%s", viper.GetString("port"), string(convertedId))
	s.repo.AddShortURL()
}
