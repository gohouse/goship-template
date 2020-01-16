package helper

import (
	"github.com/gohouse/t"
	"net/http"
)

// ApiReturn 方法之间的通信对象
type ApiReturn struct {
	Code int         `json:"code"`
	Data interface{} `json:"result"`
	Msg  interface{} `json:"msg"`
}

func NewApiReturn(code int, data, msg interface{}) ApiReturn {
	return ApiReturn{Code: code, Data: data, Msg: msg}
}

// SuccessReturn 成功返回
// SuccessReturn(),SuccessReturn("成功了"),SuccessReturn([]int{1,2},200)
func SuccessReturn(args ...interface{}) ApiReturn {
	var code = http.StatusOK
	var data interface{}
	switch len(args) {
	case 2:
		code = t.New(args[1]).Int()
		data = t.New(args[1]).Interface()
	case 1:
		data = args[0]
	}
	return NewApiReturn(code, data, "success")
}

// FailReturn 失败返回
// FailReturn(),FailReturn("操作失败"),FailReturn("已经签到过了",420)
func FailReturn(args ...interface{}) ApiReturn {
	var code = http.StatusBadRequest
	var msg = "fail"
	switch len(args) {
	case 2:
		code = t.New(args[1]).Int()
	case 1:
		msg = t.New(args[0]).String()
	}
	return NewApiReturn(code, nil, msg)
}

// QueryReturn 返回的语法糖,可以处理成功或失败的查询处理
func QueryReturn(res interface{}, err error) ApiReturn {
	if err != nil {
		return FailReturn(err.Error())
	}
	return SuccessReturn(res)
}

// ExecReturn 返回的语法糖,可以处理成功或失败的入库或更改处理
func ExecReturn(aff int64, err error) ApiReturn {
	if err != nil {
		return FailReturn(err.Error())
	}
	if aff == 0 {
		return FailReturn("操作失败")
	}
	return SuccessReturn()
}
