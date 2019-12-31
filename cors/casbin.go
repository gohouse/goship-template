package cors

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/gohouse/e"
	"github.com/gohouse/jwt"
	"net/http"
	"goshipdemo/helper"
)

func PermissionCheck(en *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		//根据上下文获取载荷claims 从claims获得role
		claims := c.MustGet("claims").(*jwt.CustomClaims)
		sub := claims.UserData["role_name"]
		//sub := t.New(claims.UserData["mobile"]).String()
		//检查权限
		obj := helper.GetParam(c, "rule")
		res, err := en.Enforce(sub, obj, "all")
		if err != nil {
			helper.ErrorBack(c, err, http.StatusUnauthorized)
			c.Abort()
			return
		}
		if res {
			c.Next()
		} else {
			helper.ErrorBack(c, e.New("权限不足"), http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
