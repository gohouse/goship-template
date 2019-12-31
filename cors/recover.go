package cors

import (
	"github.com/gin-gonic/gin"
)

func Recover()  gin.HandlerFunc {
	return func(c *gin.Context) {
		//defer func() {
		//	if r := recover(); r != nil {
		//		helper.ErrorBack(c, e.New(fmt.Sprintf("receive panic: %s", r)))
		//	}
		//}()
		c.Next()
	}
}
