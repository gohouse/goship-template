package booter

import (
	"github.com/gohouse/gopass"
	"github.com/gohouse/gopass/rules"
)

// 自定义错误信息
var msgs = gopass.Data{
	"required": "参数不能为空",
	"length":   "参数长度有误",
	"max":      "参数长度超限",
	"min":      "参数长度不足",
	"numberic": "参数必须为数字",
}

func BootValidator() *gopass.Validator {
	return gopass.NewValidator(gopass.Message(msgs)).Use(rules.Default())
}
