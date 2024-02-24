package main

import (
	"context"
	"fmt"

	"github.com/Flakannon/knuckles/internal/art"
	"github.com/Flakannon/knuckles/internal/env"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) {
	fmt.Printf("Starting Knuckles %s", env.GetAppVersion())
	fmt.Print(art.GetKnucklesArt())
}
