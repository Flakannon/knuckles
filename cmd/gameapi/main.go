package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Flakannon/knuckles/src/api"
	"github.com/Flakannon/knuckles/src/modules/env"
	"github.com/Flakannon/knuckles/src/modules/event/sqs"
)

func main() {
	sqsEventSourceConfig, err := env.LoadEventSourceConfig()
	if err != nil {
		fmt.Println("Error loading sqs event source config")
		return
	}
	sqsEventSource := sqs.NewSQSSource(sqsEventSourceConfig)

	app := api.App{
		Publisher: sqsEventSource,
	}

	engine := api.NewEngine(app)
	srv := &http.Server{
		Addr:    ":8081",
		Handler: engine,
	}

	go func() {
		log.Print("Starting server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan // block
	log.Print("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Print("Server exiting safely")
}
