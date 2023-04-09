package model

import "time"

type Application struct {
	AppId       int
	AppCode     string
	AppName     string
	Description string
	OperatedBy  string
	OperatedAt  time.Time
}
