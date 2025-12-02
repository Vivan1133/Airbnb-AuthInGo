package main

import (
	"AuthInGo/app"
	"AuthInGo/config/env"
)

func main() {

	config.LoadEnv()

	config := app.NewConfig()

	server := app.NewServer(*config)

	server.Run()

}