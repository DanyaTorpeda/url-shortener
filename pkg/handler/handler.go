package handler

import (
	"url-shortener/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/shorten", h.createShortURL)
		api.GET("/:short_url", h.redirectURL)
	}

	return router
}
