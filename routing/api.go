package routing

import (
	"github.com/gin-gonic/gin"
	"goshipdemo/api"
)

func ApiRun(route *gin.RouterGroup)  {
	route.GET("test",v1(r(api.Todo)))
}