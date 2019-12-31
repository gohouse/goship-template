package api

import (
	"github.com/gin-gonic/gin"
	"goshipdemo/helper"
	"net/http"
)

func Todo(c *gin.Context) helper.ApiReturn {
	return helper.SuccessReturn("这是 api todo 正确返回")
}

func TodoFail(c *gin.Context) helper.ApiReturn {
	return helper.FailReturn("这是 api 失败返回")
}
func TodoGin(c *gin.Context) {
	c.String(http.StatusOK,"这是 api gin 直接输出 ...")
}