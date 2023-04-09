package application

import (
	"github.com/xuanye/keystone-go/application/dto"
	"github.com/xuanye/keystone-go/domain/service"
)

type Handler struct {
	service *service.ApplicationService
}

func NewHandler(s *service.ApplicationService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) FindAll() []dto.DtoApplication {
	apps := h.service.FindAll()

	var list []dto.DtoApplication
	for _, app := range apps {
		list = append(list, dto.DtoApplication{
			AppId:       app.AppId,
			AppName:     app.AppName,
			AppCode:     app.AppCode,
			Description: app.Description,
			OperatedAt:  app.OperatedAt.Unix(),
			OperatedBy:  app.OperatedBy,
		})
	}
	return list
}
