package dto

import "time"

type DtoApplication struct {
	AppId       int       `json:"appId"`
	AppName     string    `json:"appName"`
	AppCode     string    `json:"appCode"`
	Description string    `json:"description"`
	OperatedAt  time.Time `json:"operatedAt"`
	OperatedBy  string    `json:"operatedBy"`
}
