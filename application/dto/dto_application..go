package dto

type DtoApplication struct {
	AppId       int    `json:"appId"`
	AppName     string `json:"appName"`
	AppCode     string `json:"appCode"`
	Description string `json:"description"`
	OperatedAt  int64  `json:"operatedAt"`
	OperatedBy  string `json:"operatedBy"`
}
