package handler

import (
	"net/http"
	shortener "url-shortener"

	"github.com/gin-gonic/gin"
)

type ReceivedURL struct {
	LongURL string `json:"long_url"`
}

func (h *Handler) createShortURL(c *gin.Context) {
	var input ReceivedURL

	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	//TODO add validation
	url := shortener.URL{LongURL: input.LongURL}
	shortURL, err := h.service.Shortener.CreateShortURL(url)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newAPIResponse("success", map[string]string{
		"short_url": shortURL,
	}))
}

func (h *Handler) redirectURL(c *gin.Context) {
	shortURL := c.Param("short_url")
	if shortURL == "" {
		newErrorMessage(c, http.StatusBadRequest, "invalid short url data")
		return
	}

	longURL, err := h.service.Shortener.GetLongURL(shortURL)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusMovedPermanently, longURL)
}
