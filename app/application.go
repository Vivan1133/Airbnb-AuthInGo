package app

import (
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

func NewConfig(addr string) *Config {
	return &Config{ Addr: addr}
}


type Server struct {
	Config Config
}

func NewServer(config Config) *Server {
	return &Server{Config: config}
}

func (server *Server) Run() error {

	s := &http.Server {
		Addr: server.Config.Addr,
		Handler: nil,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server running @ port: ",server.Config.Addr)

	return s.ListenAndServe()

}


