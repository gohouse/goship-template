package routing

import (
	"github.com/gin-gonic/gin"
	"goshipdemo/cors"
)

func v1(arg gin.HandlerFunc) gin.HandlerFunc {
	return cors.Version("v1",arg)
}
func v2(arg gin.HandlerFunc) gin.HandlerFunc {
	return cors.Version("v2",arg)
}
func v3(arg gin.HandlerFunc) gin.HandlerFunc {
	return cors.Version("v3",arg)
}

func r(hrh cors.HttpResponseHandler) gin.HandlerFunc {
	return cors.HttpResponse(hrh)
}
