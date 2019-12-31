package cors

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/e"
	"github.com/gohouse/jwt"
	"net/http"
	"goshipdemo/helper"
)

// JWTAuth 中间件，检查token
func JWTAuth(j *jwt.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			helper.ErrorBack(c, e.New("请求未携带token，无权限访问"))
			c.Abort()
			return
		}

		//log.Print("get token: ", token)

		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				c.JSON(http.StatusOK, helper.GinBack{
					Code: http.StatusUnauthorized,
					Msg:  "授权已过期",
					Data: nil,
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, helper.GinBack{
				Code: http.StatusUnauthorized,
				Msg:  err.Error(),
				Data: nil,
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		c.Next()
	}
}
