package booter

import (
	"github.com/gin-gonic/gin"
)
func BootGin() *gin.Engine {
	r := gin.Default()
	return r
}


