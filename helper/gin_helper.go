package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gohouse/e"
	"github.com/gohouse/gopass"
	"github.com/gohouse/jwt"
	"github.com/gohouse/t"
	"net/http"
	"strings"
)

// GinBack gin返回统一结构封装
type GinBack struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"fail"`
	Data interface{} `json:"data"`
}

// ErrorBack gin框架的错误返回封装
func ErrorBack(c *gin.Context, err interface{}, statusCode ...int) {
	var msg string
	if v, ok := err.(e.Error); ok {
		msg = v.Error()
		// 记录错误堆栈信息,方便调试
		//log.Println(v.ErrorWithStack())
	} else {
		v := err.(error)
		msg = v.Error()
	}
	if msg == "" {
		msg = "fail"
	}
	code := http.StatusBadRequest
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	c.JSON(http.StatusOK, GinBack{
		Msg:  msg,
		Code: code,
	})
	c.Abort()
}

// OkBack gin框架的成功返回封装
func OkBack(c *gin.Context, data interface{}, statusCode ...int) {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	c.JSON(http.StatusOK, GinBack{
		Msg:  "success",
		Data: data,
		Code: code,
	})
}

func GetJwtClaims(c *gin.Context) interface{} {
	value, exists := c.Get("claims")
	if !exists {
		return nil
	}
	return value
}

func GetJwtClaimsItem(c *gin.Context, item string) interface{} {
	if claims := GetJwtClaims(c); claims != nil {
		return claims.(*jwt.CustomClaims).UserData[item]
	}
	return nil
}

func BuildPageParams(c *gin.Context) (limit, page int) {
	limit = t.New(c.DefaultQuery("limit", "20")).Int()
	page = t.New(c.DefaultQuery("page", "1")).Int()
	if page < 1 {
		page = 1
	}
	return
}

// GetParam 获取gin的某一个请求参数
func GetParam(c *gin.Context, field string) string {
	switch strings.ToLower(c.Request.Method) {
	case "get", "delete":
		if res, ok := c.GetQuery(field); ok {
			return res
		}
	case "post", "put":
		if res, ok := c.GetPostForm(field); ok {
			return res
		}
	default:
		if res := c.Param(field); res != "" {
			return res
		}
	}
	return ""
}

func BuildWhere(c *gin.Context, fields []string) (map[string]interface{}, int) {
	var where = MapData{}
	var dataNum int

	for _, item := range fields {
		switch c.Request.Method {
		case "GET", "DELETE":
			if res, ok := c.GetQuery(item); ok && res != "" {
				where[item] = res
				dataNum++
			}
		case "POST", "PUT":
			if res, ok := c.GetPostForm(item); ok && res != "" {
				where[item] = res
				dataNum++
			}
		}

	}
	return where, dataNum
}

func BuildWhereMulti(c *gin.Context, fields []string) ([][]interface{}, int) {
	var where = [][]interface{}{}
	var dataNum int

	for _, item := range fields {
		switch c.Request.Method {
		case "GET", "DELETE":
			if res, ok := c.GetQuery(item); ok && res != "" {
				where = append(where, []interface{}{item, res})
				dataNum++
			}
		case "POST", "PUT":
			if res, ok := c.GetPostForm(item); ok && res != "" {
				where = append(where, []interface{}{item, res})
				dataNum++
			}
		}

	}
	return where, dataNum
}

func BuildValidateData(c *gin.Context, rule gopass.Rules) map[string]interface{} {
	var params = make(map[string]interface{})
	for k, _ := range rule {
		switch c.Request.Method {
		case "GET", "DELETE":
			if v, ok := c.GetQuery(k); ok {
				params[k] = v
			}
		case "POST", "PUT":
			if v, ok := c.GetPostForm(k); ok {
				params[k] = v
			}
		}
	}
	return params
}

func HandleFinish(c *gin.Context, res interface{}, err error) {
	if err != nil {
		ErrorBack(c, err)
		return
	}
	OkBack(c, res)
}

func HandleExec(aff int64, err error) e.Error {
	if err != nil {
		return e.New(err.Error())
	}
	if aff == 0 {
		return e.New("入库失败")
	}
	return nil
}

func HandleExecFinish(c *gin.Context, aff int64, err error) {
	HandleFinish(c, nil, HandleExec(aff, err))
}

func BuildLike(c *gin.Context, where *[][]interface{}, field string) {
	switch c.Request.Method {
	case "GET", "DELETE":
		if r, ok := c.GetQuery(field); ok {
			*where = append(*where, []interface{}{field, "like", fmt.Sprint("%", r, "%")})
		}
	case "POST", "PUT":
		if r, ok := c.GetPostForm(field); ok {
			*where = append(*where, []interface{}{field, "like", fmt.Sprint("%", r, "%")})
		}
	}
}
