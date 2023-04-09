package service

import (
	"github.com/xuanye/keystone-go/domain/model"
	"time"
)

type ApplicationService struct {
}

func NewApplicationService() *ApplicationService {
	return &ApplicationService{}
}

func (service *ApplicationService) FindAll() []model.Application {

	return []model.Application{
		{
			AppId:       1,
			AppName:     "Keystone",
			AppCode:     "Keystone",
			Description: "Keystone is an application",
			OperatedBy:  "admin",
			OperatedAt:  time.Now(),
		},
	}
}
