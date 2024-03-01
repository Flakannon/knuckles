package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Flakannon/knuckles/src/api"
)

func main() {

	engine := api.NewEngine()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		log.Print("Starting server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// helps to gracefully shutdown the server and give the right signal especially useful for ecs if the api is long lived
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // block until signal received
	log.Print("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Print("Server exiting safely")

}
