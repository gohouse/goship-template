package util

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gopass"
	"github.com/gohouse/gorose/v2"
	"github.com/gohouse/jwt"
	"goshipdemo/booter"
	"goshipdemo/helper"
	"github.com/gohouse/gosms"
)

func DB() gorose.IOrm {
	return booter.NewBooter().Engin.NewOrm()
}

func GetConfig(key string) interface{} {
	return booter.NewBooter().AppConfig.GetConfig(key)
}

func Validate(c *gin.Context, rules gopass.Rules, msgs ...gopass.Data) error {
	return booter.NewBooter().Validator.
		Validate( helper.BuildValidateData(c, rules),rules,msgs...)
}

func Jwt() *jwt.JWT {
	return booter.NewBooter().JWT
}

func Casbin() *casbin.Enforcer {
	return booter.NewBooter().Enforcer
}

func Enforce(rvals ...interface{}) (bool, error) {
	return Casbin().Enforce(rvals...)
}

func GoSMS() *gosms.GoSMS {
	return booter.NewBooter().GoSMS
}

func BaseConfig() *booter.BaseConfig {
	return booter.NewBooter().BaseConfig
}

func OSS() *booter.OSS {
	return booter.NewBooter().OSS
}