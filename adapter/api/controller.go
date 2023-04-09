package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Controller interface {
	Build(engine *gin.RouterGroup)
}

type ControllerGroup struct {
	dig.In

	Controllers []Controller `group:"controller"`
}
