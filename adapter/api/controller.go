package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
)

type Controller interface {
	Build(engine *gin.RouterGroup)
}

type ControllerGroup struct {
	dig.In

	Controllers []Controller `group:"controller"`
}

type BaseController struct {
	context *gin.Context
}

func (bc *BaseController) WithContext(c *gin.Context) *BaseController {
	bc.context = c

	return bc
}

func (bc *BaseController) Ok(data interface{}, message string) {
	bc.context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
		"data":    data,
	})
}

func (bc *BaseController) Error(code int, message string) {
	// note:MapStatusCode or not
	bc.context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}
