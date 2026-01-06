package models

import "time"

type Role struct {
	Id          int
	Name        string
	Description string
	Created_at  time.Time
	Updated_at  time.Time
}

type Permission struct {
	Id          int
	Name        string
	Description string
	Resource    string
	Action      string
	Created_at  string
	Updated_at  string
}