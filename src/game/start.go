package game

import (
	"context"
	"log"
	"math/rand"
	"time"

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
	board := NewBoard()
	randomiser := rand.New(rand.NewSource(time.Now().UnixNano()))

	shipSize := 3
	board.PlaceShip(shipSize, randomiser)
	ship2Size := 4
	board.PlaceShip(ship2Size, randomiser)

	board.PrintBoard()
	return "started game"
}
