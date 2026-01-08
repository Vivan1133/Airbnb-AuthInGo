package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	dbRep "AuthInGo/db/repositories"
	dbRoleRep "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
	dbConfig "AuthInGo/config/db"
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

	db, err := dbConfig.SetupDB()

	if err != nil {
		fmt.Println("Error connecting to db")
		return err
	}

	urep := dbRep.NewUserRepository(db)
	rrep := dbRoleRep.NewRoleRepository(db)
	rrprep := dbRoleRep.NewRolesPermissions(db)
	rprep := dbRoleRep.NewPermissionsRepository(db)
	urrep := dbRep.NewUserRoles(db)


	us := services.NewUserService(urep)
	rs := services.NewRoleService(rrep, rrprep, rprep, urrep)

	uc := controllers.NewUserController(us)
	rc := controllers.NewRoleController(rs)

	urou := router.NewUserRouter(uc)
	rrou := router.NewRoleRouter(rc)
	

	s := &http.Server {
		Addr: server.Config.Addr,
		Handler: router.CreateRouter(urou, rrou),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server running @ port: ",server.Config.Addr)

	return s.ListenAndServe()

}


