package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanye/keystone-go/application/feature/application"
)

type ApplicationController struct {
	handler *application.Handler
	BaseController
}

func NewApplicationController(h *application.Handler) Controller {
	return &ApplicationController{handler: h}
}
func (c *ApplicationController) Build(g *gin.RouterGroup) {
	g.GET("/apps", c.findAll)
}

func (c *ApplicationController) findAll(ctx *gin.Context) {

	data := c.handler.FindAll()
	c.WithContext(ctx).Ok(data, "")

}
