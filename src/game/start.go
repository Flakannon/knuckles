package game

import "context"

func StartGame(ctx context.Context) {
	// publish game start message to sns that will fan out and trigger two battle lambdas to start

}

