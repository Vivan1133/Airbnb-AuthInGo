package main

import (
	"AuthInGo/app"
)

func main() {

	config := app.NewConfig(":3004")

	server := app.NewServer(*config)

	server.Run()

}