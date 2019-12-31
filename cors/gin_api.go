package cors

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/t"
	"net/http"
	"goshipdemo/helper"
)

// Version 版本控制
// route.Group("/api/v3")
// route.GET("/getUserInfo", Version("v2",api.Test2)), api.Test)
// 如果请求的是v3版本,因为最新版本是v2,所以会自动选择v2版本
// 如果请求的是v1版本,因为比v2早,则会自动选择最后一个,即 api.Test
func Version(version string, handler gin.HandlerFunc) gin.HandlerFunc {
	versionNo := t.New(version[1:]).Int()
	return func(c *gin.Context) {
		versionRequestNo := t.New(c.Param("version")[1:]).Int()
		//判断版本号...
		if versionRequestNo >= versionNo {
			handler(c)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// 统一返回
// route.GET("/test", Version("v2",HttpResponse(api.Test4)), Version("v1",HttpResponse(api.Test3)))
// 这个函数的目的是为了统一返回,同时,也可以方便互相调用和控制
type HttpResponseHandler func(c *gin.Context) helper.ApiReturn
func HttpResponse(hrh HttpResponseHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := hrh(c)
		if res.Code == 200 {
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, res)
		}
	}
}

type HttpResponseHandler2 func(c *gin.Context) (res interface{}, err error)
func HttpResponse2(hrh HttpResponseHandler2) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code interface{}
		res, err := hrh(c)
		if err != nil {
			code = 400
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"data": res,
				"msg":  err.Error(),
			})
		} else {
			code = 200
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"data": res,
				"msg":  "fail",
			})
		}
	}
}
