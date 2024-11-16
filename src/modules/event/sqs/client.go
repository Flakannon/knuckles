package sqs

import (
	"context"
	"fmt"
	"log"

	"github.com/Flakannon/knuckles/src/modules/env"
	"github.com/Flakannon/knuckles/src/modules/event"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SQSClient struct {
	Source event.EventClient
	client *sqs.Client
	ctx    context.Context
}

func NewSQSSource(config env.EventSourceConfig) (*SQSClient, error) {
	cfg, err := awsconfig.LoadDefaultConfig(context.Background(), awsconfig.WithRegion("eu-west-2"))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS configuration: %w", err)
	}

	client := sqsClientWithResolution(cfg)

	return &SQSClient{
		Source: event.EventClient{
			Config: config,
		},
		client: client,
		ctx:    context.Background(), // will pas in
	}, nil
}

func sqsClientWithResolution(cfg aws.Config) *sqs.Client {
	return sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		localstackURL := env.GetLocalStackURL()
		if localstackURL != "" {
			o.BaseEndpoint = aws.String("http://" + localstackURL + ":4566")
		}
	})
}

func (s *SQSClient) Publish(message string) error {
	_, err := s.client.SendMessage(s.ctx, &sqs.SendMessageInput{
		QueueUrl:    &s.Source.Config.URL,
		MessageBody: aws.String(message),
	})
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("Published message to SQS queue %s: %s", s.Source.Config.URL, message)
	return nil
}

func (s *SQSClient) Subscribe(queueName string, processMessage func(message types.Message) error) error {
	queueURL, err := s.GetQueueUrl(queueName)
	if err != nil {
		return fmt.Errorf("failed to get queue URL: %w", err)
	}

	for {
		msgOutput, err := s.client.ReceiveMessage(s.ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            &queueURL,
			MaxNumberOfMessages: 10,
			WaitTimeSeconds:     20,
		})
		if err != nil {
			return fmt.Errorf("failed to receive messages: %w", err)
		}

		for _, msg := range msgOutput.Messages {
			err := processMessage(msg)
			if err != nil {
				log.Printf("Failed to process message: %s", err)
				continue
			}

			_, err = s.client.DeleteMessage(s.ctx, &sqs.DeleteMessageInput{
				QueueUrl:      &queueURL,
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				log.Printf("Failed to delete message: %s", err)
			}
		}
	}
}

func (s *SQSClient) GetQueueUrl(queueName string) (string, error) {
	urlResult, err := s.client.GetQueueUrl(context.Background(), &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		return "", fmt.Errorf("failed to resolve queue URL: %w", err)
	}
	return *urlResult.QueueUrl, nil
}
