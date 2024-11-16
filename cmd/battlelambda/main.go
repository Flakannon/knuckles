package main

import (
	"context"
	"fmt"

	"github.com/Flakannon/knuckles/src/modules/art"
	"github.com/Flakannon/knuckles/src/modules/env"
	"github.com/Flakannon/knuckles/src/modules/event/sqs"
	"github.com/Flakannon/knuckles/src/publisher"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) {
	fmt.Printf("Starting Knuckles %s", env.GetAppVersion())
	fmt.Println(art.GetKnucklesArt())

	// kafka source
	// kafkaEventSourceConfig, err := env.LoadEventSourceConfig()
	// if err != nil {
	// 	fmt.Println("Error loading kafka event source config")
	// 	return
	// }

	// kafkaEventSource := kafka.NewKafkaSource(kafkaEventSourceConfig)
	// publisher.PublishMessage(kafkaEventSource, "Hello, Kafka!")

	sqsEventSourceConfig, err := env.LoadEventSourceConfig()
	if err != nil {
		fmt.Println("Error loading sqs event source config")
		return
	}

	sqsEventSource, err := sqs.NewSQSSource(sqsEventSourceConfig)
	if err != nil {
		fmt.Println("Error creating sqs event source")
		return
	}
	publisher.PublishMessage(sqsEventSource, "Hello, SQS!")
}
