package config

import (
	env "AuthInGo/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = env.GetString("DB_USER", "root")
	cfg.Addr = env.GetString("DB_ADDR", "127.0.0.1:3004")
	cfg.DBName = env.GetString("DB_NAME", "airbnb_auth_dev")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Passwd = env.GetString("DB_PASSWORD", "adminadmin")

	fmt.Println(cfg.FormatDSN())

	db, err := sql.Open("mysql", cfg.FormatDSN())
    
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	pingErr := db.Ping()
    if pingErr != nil {
		fmt.Println(pingErr)
		return nil, pingErr
    }
    fmt.Println("Connected!")

	return db, nil

}