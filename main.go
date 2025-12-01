package main

import (
	"AuthInGo/app"
	"AuthInGo/config/env"
	dbConfig "AuthInGo/config/db"
)

func main() {

	config.LoadEnv()

	config := app.NewConfig()

	server := app.NewServer(*config)

	dbConfig.SetupDB()

	server.Run()

}