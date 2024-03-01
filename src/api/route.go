package api

import (
	"github.com/gin-gonic/gin"
)

func healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleHealthCheck(c)
	}
}
