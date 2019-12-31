package booter

import (
	"github.com/gohouse/jwt"
)

func BootJWT(baseConfig *BaseConfig) *jwt.JWT {
	var jwtOptions = jwt.Options{
		Secret: baseConfig.Jwt.Secret,
		Expire: baseConfig.Jwt.Expire,
	}
	if jwtOptions.Secret == "" {
		panic("jwt缺少secret配置项")
	}
	return jwt.NewJWT(&jwtOptions)
}
