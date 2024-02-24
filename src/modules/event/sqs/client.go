package sqs

import (
	"log"

	"github.com/Flakannon/knuckles/src/modules/env"
	"github.com/Flakannon/knuckles/src/modules/event"
)

type SQSClient struct {
	Source event.EventClient
}

func (s SQSClient) Publish(message string) error {
	log.Print("Publishing message to SQS: ", message)
	return nil
}

func (s SQSClient) Subscribe() error {
	log.Print("Subscribing to SQS")
	return nil
}

func NewSQSSource(config env.EventSourceConfig) SQSClient {
	return SQSClient{
		Source: event.EventClient{
			Config: config,
		},
	}
}
