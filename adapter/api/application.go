package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanye/keystone-go/application/feature/application"
)

type ApplicationController struct {
	handler *application.Handler
}

func NewApplicationController(h *application.Handler) Controller {
	return &ApplicationController{handler: h}
}
func (controller *ApplicationController) Build(g *gin.RouterGroup) {
	g.GET("/apps", controller.findAll)
}

func (controller *ApplicationController) findAll(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"code":    0,
		"message": "",
		"data":    controller.handler.FindAll(),
	})
}
