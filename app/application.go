package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	db "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

func NewConfig() *Config {

	addr := config.GetString("PORT", ":3005")

	return &Config{ Addr: addr}
}


type Server struct {
	Config 	Config
}

func NewServer(config Config) *Server {
	return &Server{Config: config,}
}

func (server *Server) Run() error {

	urep := db.NewUserRepository()

	us := services.NewUserService(urep)

	uc := controllers.NewUserController(us)

	urou := router.NewUserRouter(uc)
	

	s := &http.Server {
		Addr: server.Config.Addr,
		Handler: router.CreateRouter(urou),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server running @ port: ",server.Config.Addr)

	return s.ListenAndServe()

}


