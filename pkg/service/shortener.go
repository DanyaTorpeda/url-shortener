package service

import (
	"fmt"
	"strconv"
	shortener "url-shortener"
	"url-shortener/pkg/repository"

	"github.com/jxskiss/base62"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ShortenerService struct {
	repo repository.Shortener
}

func NewShortenerService(repo repository.Shortener) *ShortenerService {
	return &ShortenerService{repo: repo}
}

func (s *ShortenerService) CreateShortURL(input shortener.URL) (string, error) {
	id, err := s.repo.CreateLongURL(input)
	if err != nil {
		return "", err
	}

	shortURL := encodeID(id)
	logrus.Print(shortURL)
	//shortURL := fmt.Sprintf("http://localhost:%s/%s", viper.GetString("port"), convertedId)
	err = s.repo.AddShortURL(id, shortURL)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://localhost:%s/%s", viper.GetString("port"), shortURL), nil
}

func (s *ShortenerService) GetLongURL(shortURL string) (string, error) {
	id, err := decodeID(shortURL)
	if err != nil {
		return "", err
	}

	longURL, err := s.repo.GetLongURL(id)
	if err != nil {
		return "", err
	}

	return longURL, nil
}

func encodeID(id int) string {
	return base62.EncodeToString([]byte(strconv.Itoa(id)))
}

func decodeID(shortURL string) (int, error) {
	val, err := base62.DecodeString(shortURL)
	if err != nil {
		return 0, err
	}

	id, err := strconv.Atoi(string(val))
	if err != nil {
		return 0, err
	}

	return id, nil
}
