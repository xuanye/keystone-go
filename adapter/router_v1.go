package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanye/keystone-go/adapter/api"
)

func NewRouterV1(g api.ControllerGroup) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	for _, c := range g.Controllers {
		c.Build(v1)
	}

	return r

}
