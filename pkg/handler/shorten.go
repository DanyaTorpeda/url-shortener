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
	h.service.Shortener.CreateShortURL(url)
}

func (h *Handler) redirectURL(c *gin.Context) {

}
