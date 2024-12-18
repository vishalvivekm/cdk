package main

import (
	"fmt"
	"vishalvivekm/lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

func HandleRequest(event MyEvent) (string, error){
	if event.Username == "" {
		return "",fmt.Errorf("username can't be empty")
	}
	return fmt.Sprintf("Successfully called by ~ %s", event.Username), nil
}

func main() {
	// lambda.Start is not calling this func, but is passed a func to call, when the lambda
	// is invoked
    myApp := app.NewApp()
	lambda.Start(myApp.ApiHandler.RegisterUserHandler)
}