package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func newErrorMessage(c *gin.Context, code int, message string) {
	logrus.WithFields(logrus.Fields{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"status": code,
	}).Error(message)
	c.AbortWithStatusJSON(code, errorMessage{Message: message, Code: code})
}

type APIResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func newAPIResponse(status string, data interface{}) *APIResponse {
	return &APIResponse{Status: status, Data: data}
}
