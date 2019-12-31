package booter

import (
	"github.com/gohouse/gorose/v2"
	"github.com/gohouse/gosms"
	"github.com/gohouse/gosms/adapter/drivers"
	"github.com/gohouse/gosms/adapter/sdks"
)

func BootGoSMS(baseConfig *BaseConfig, db *gorose.Engin) *gosms.GoSMS {
	return gosms.NewGoSMS(db, drivers.NewMysqlDriver(), gosms.Sdk{
		China:  sdks.NewAliyunSdk(baseConfig.AliyunOptions),
		Global: sdks.NewTwilioSdk(baseConfig.TwilioOptions),
	})
}
