package api

import (
	"github.com/Flakannon/knuckles/src/game"
	"github.com/gin-gonic/gin"
)

func healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleHealthCheck(c)
	}
}

func startGame(app App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var gameConfig game.GameConfig
		if err := c.ShouldBind(&gameConfig); err != nil {
			c.JSON(400, gin.H{
				"message": "missing required fields",
			})
			return
		}
		HandleStartGame(c, app, gameConfig)
	}
}
