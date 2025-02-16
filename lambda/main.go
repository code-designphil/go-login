package main

import (
	"fmt"
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleRequest(event Event) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by %s", event.Username), nil
}

func main() {
	app := app.NewApp()
	lambda.Start(app.ApiHandler.RegisterUserHandler)
}
