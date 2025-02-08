package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vishalvivekm/lambda-func/database"
	"vishalvivekm/lambda-func/types"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler{
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var registerUser types.RegisterUser

	err := json.Unmarshal([]byte(request.Body), &registerUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: "Invalid Request",
			StatusCode: http.StatusBadRequest,
		}, err
	}
	if registerUser.Username == "" || registerUser.Password == "" {
		return events.APIGatewayProxyResponse{
			Body: "Invalid request - fields empty",
			StatusCode: http.StatusBadRequest,
		}, err
	}
	userExists, err := api.dbStore.DoesUserExist(registerUser.Username)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("there was an error checking if user exists %w", err)
	}
	
	if userExists {
		return events.APIGatewayProxyResponse{
			Body: "User already exists",
			StatusCode: http.StatusConflict,
		}, fmt.Errorf("user with the username already exists")
	}
	user, err := types.NewUser(registerUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("could not create new user %w", err)
	}
	err = api.dbStore.InsertUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("errro inserting the user %w", err)
	}

	return events.APIGatewayProxyResponse{
		Body: "Successfully registered user",
		StatusCode: http.StatusOK,
	}, nil
}