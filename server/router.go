package main

import (
	"github.com/gin-gonic/gin"

	"github.com/SergeyShpak/propro/server/handlers"
)

func newRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handlers.Ping())
	return r
}
