package main

import (
	"fmt"
	"vishalvivekm/lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"vishalvivekm/lambda-func/middleware"
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


func ProtectedHanlder(request events.APIGatewayProxyRequest)(events.APIGatewayProxyResponse, error){
	return events.APIGatewayProxyResponse{
		Body: "a secret resource",
		StatusCode: http.StatusOK,
	}, nil
}
func main() {
	// lambda.Start is not calling this func, but is passed a func to call, when the lambda
	// is invoked
    myApp := app.NewApp()
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.Path {
		case "/register":
			return myApp.ApiHandler.RegisterUserHandler(request)
		case "/login":
			return myApp.ApiHandler.LoginUser(request)
		case "/protected":
			return middleware.ValidateJWTMiddleware(ProtectedHanlder)(request)
		default:
			return events.APIGatewayProxyResponse{
			Body: "Not found",
			StatusCode: http.StatusNotFound,
		}, nil

		}
	})
}