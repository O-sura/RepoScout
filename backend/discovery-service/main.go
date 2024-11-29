package main

import (
	"net/http"

	"example.com/discovery-service/api"
	"example.com/discovery-service/utils"
)

func main() {
	router := api.SetupRouter()
	logger := utils.NewLogger()

	logger.Info("Repo discovery-service on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Critical("Could not start server: " +  err.Error())
	}
}