package api

import (
	"github.com/Flakannon/knuckles/src/game"
	"github.com/gin-gonic/gin"
)

func HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "healthy",
	})
}

func HandleStartGame(ctx *gin.Context, app App, config game.GameConfig) {
	result := game.StartGame(ctx, app.Publisher, config)
	if result == "" {
		ctx.JSON(400, gin.H{
			"message": "game start issue",
		})
	}
	ctx.JSON(200, result)
}
