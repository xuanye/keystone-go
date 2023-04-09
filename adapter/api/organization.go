package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanye/keystone-go/application/feature/application"
)

type OrganizationController struct {
	handler *application.Handler
}

func NewOrganizationController(h *application.Handler) Controller {
	return &OrganizationController{handler: h}
}
func (controller *OrganizationController) Build(g *gin.RouterGroup) {
	g.GET("/orgs", controller.findAll)
}

func (controller *OrganizationController) findAll(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"code":    0,
		"message": "",
		"data":    controller.handler.FindAll(),
	})
}
