package main

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/spie/gospeed/api"
	"github.com/spie/gospeed/db"
	"github.com/spie/gospeed/speedtest"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	connectionHandler := createConnectionHandler()
	defer connectionHandler.Close()

	speedtestRepo := createSpeedtestRepositoryAndMigrate(connectionHandler)
	client := createSpeedtestClient()

	finished := speedtest.NewHandler(client, speedtestRepo).Run()

	api.Run(
		os.Getenv("HOST_ADDRESS"),
		os.Getenv("HOST_PORT"),
		speedtest.NewController(speedtestRepo),
	)

	<-finished
}

func createConnectionHandler() db.ConnectionHandler {
	connectionHandler, err := db.NewConnectionHandler(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	if err != nil {
		panic(err)
	}

	return connectionHandler
}

func createSpeedtestRepositoryAndMigrate(connectionHandler db.ConnectionHandler) speedtest.Repository {
	speedtestRepo := speedtest.NewRepository(connectionHandler)
	speedtestRepo.AutoMigrate()

	return speedtestRepo
}

func createSpeedtestClient() speedtest.Client {
	client, err := speedtest.NewClient()
	if err != nil {
		panic(err)
	}

	return client
}
