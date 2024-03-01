package aws

import (
	"context"
	"log"

	"github.com/Flakannon/knuckles/src/modules/env"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func customResolver() aws.EndpointResolverWithOptionsFunc {
	localStackURL := env.GetLocalStackURL()
	return func(service string, region string, options ...interface{}) (aws.Endpoint, error) {
		if localStackURL != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           localStackURL,
				SigningRegion: "eu-west-2",
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	}
}

func GetConfig(ctx context.Context) (aws.Config, error) {
	options := config.WithEndpointResolverWithOptions(customResolver())
	cfg, err := config.LoadDefaultConfig(ctx, options)
	if err != nil {
		log.Print("[AWS] Error loading AWS config: ", err)
		return aws.Config{}, err //TODO make that a custom error with a more friendly message for the user
	}
	return cfg, nil
}
