package api

import (
	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	g := gin.Default()

	g.GET("/health", healthCheck())

	return g
}
