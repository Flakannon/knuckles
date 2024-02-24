package kafka

import (
	"log"

	"github.com/Flakannon/knuckles/src/modules/env"
	"github.com/Flakannon/knuckles/src/modules/event"
)

type KafkaClient struct {
	Source event.EventClient
}

func (k KafkaClient) Publish(message string) error {
	log.Printf("Publishing message to Kafka: %s", message)
	return nil
}

func (k KafkaClient) Subscribe() error {
	log.Printf("Subscribing to Kafka")
	return nil
}

func NewKafkaSource(config env.EventSourceConfig) KafkaClient {
	return KafkaClient{
		Source: event.EventClient{
			Config: config,
		},
	}
}
