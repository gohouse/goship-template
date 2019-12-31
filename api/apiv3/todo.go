package apiv3

import (
	"github.com/gin-gonic/gin"
	"goshipdemo/helper"
	"net/http"
)

func Todo(c *gin.Context) helper.ApiReturn {
	return helper.SuccessReturn("这是 apiv3 todo 正确返回")
}

func TodoFail(c *gin.Context) helper.ApiReturn {
	return helper.FailReturn("这是 apiv3 失败返回")
}
func TodoGin(c *gin.Context) {
	c.String(http.StatusOK,"这是 apiv3 gin 直接输出 ...")
}