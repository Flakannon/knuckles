package api

import (
	"github.com/Flakannon/knuckles/src/publisher"
	"github.com/gin-gonic/gin"
)

type App struct {
	Publisher publisher.IPublisher
}

func NewEngine(app App) *gin.Engine {
	g := gin.Default()

	g.GET("/health", healthCheck())

	g.POST("/game/start", startGame(app))

	return g
}
