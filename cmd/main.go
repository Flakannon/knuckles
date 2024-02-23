package main

import (
	"context"
	"fmt"

	"github.com/Flakannon/knuckles/internal/ascii"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	//////////////////////////
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) {
	fmt.Println("Starting Knuckles 0.0.1")
	fmt.Print(ascii.GetKnucklesArt())
}
