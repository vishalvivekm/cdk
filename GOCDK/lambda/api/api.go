package api

import (
	"fmt"
	"vishalvivekm/lambda-func/database"
	"vishalvivekm/lambda-func/types"
	
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler{
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty parameters")
	}
	userExists, err := api.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("there was an error checking if user exists %w", err)
	}
	
	if userExists {
		return fmt.Errorf("user with the username already exists")
	}
	
	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("errro registering the user %w", err)
	}
	return nil
}