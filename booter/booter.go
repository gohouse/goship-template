package booter

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gopass"
	"github.com/gohouse/gorose/v2"
	"github.com/gohouse/gosms"
	"github.com/gohouse/jwt"
	"sync"
)

type Booter struct {
	*BaseConfig
	*gorose.Engin
	*gin.Engine
	*gopass.Validator
	*jwt.JWT
	*AppConfig
	*casbin.Enforcer
	*gosms.GoSMS
	*OSS
}

var once sync.Once
var boot *Booter

func NewBooter(files ...string) *Booter {
	once.Do(func() {
		if len(files)==0 {
			panic("请传入配置文件")
		}
		boot = &Booter{}

		// 初始化配置文件
		baseConfig := LoadBaseConfig(files[0])
		boot.BaseConfig = baseConfig

		// 数据库orm
		boot.Engin = BootGorose(baseConfig)
		// 路由
		boot.Engine = BootGin()
		// 验证器
		boot.Validator = BootValidator()
		// jwt 认证
		boot.JWT = BootJWT(baseConfig)
		// 配置
		boot.AppConfig = BootConfig(boot.Engin)
		// casbin 权限管理
		boot.Enforcer = BootCasbin(baseConfig, boot.Engin)
		// gosms 短信验证码
		boot.GoSMS = BootGoSMS(baseConfig,boot.Engin)
		// oss
		boot.OSS = BootOSS(baseConfig)

		// 驱动错误包
		BootError(baseConfig.Error)
	})

	return boot
}
