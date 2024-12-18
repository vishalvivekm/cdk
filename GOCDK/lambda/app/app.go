package app

import (
	"vishalvivekm/lambda-func/api"
	"vishalvivekm/lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App{
	db := database.NewDynamoDBClient()

	apiHandler := api.NewApiHandler(db)

	return App{
		ApiHandler: apiHandler,
	}
}