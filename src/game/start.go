package game

import (
	"context"
	"log"

	"github.com/Flakannon/knuckles/src/publisher"
)

type GameConfig struct {
	Character string `form:"character" binding:"required"`
	BoardSize string `form:"boardSize" binding:"required"`
}

func StartGame(ctx context.Context, pub publisher.IPublisher, config GameConfig) string {
	log.Print("starting game")
	log.Print("character: ", config.Character)
	log.Print("boardSize: ", config.BoardSize)

	pub.Publish("game-started")
	return "started game"
}
