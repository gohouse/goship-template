package routing

import (
	"github.com/gin-gonic/gin"
	"goshipdemo/api"
	"goshipdemo/api/apiv2"
	"goshipdemo/api/apiv3"
	"net/http"
)

func Run(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":200,
			"msg":"test",
			"data":"goship是一个根据表结构自动化生成 restful api 的框架",
		})
	})

	// 定义app请求根路由
	app := route.Group("/api/:version")

	// api 版本控制测试,分别访问
	// /api/v3/todo
	// /api/v2/todo
	// /api/v1/todo
	app.GET("/todo", v3(r(apiv3.Todo)), v2(r(apiv2.TodoFail)), api.TodoGin)

	// jwt认证
	//app.Use(cors.JWTAuth(util.Jwt()))

	ApiRun(app)
}
