package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
}
