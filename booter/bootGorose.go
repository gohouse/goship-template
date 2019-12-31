package booter

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
)

// BootGorose 驱动 gorose orm
func BootGorose(baseConfig *BaseConfig) *gorose.Engin {
	var dbc = baseConfig.Database
	var config = gorose.Config{
		Driver:          dbc.Driver,
		Dsn:             dbc.Dsn,
		SetMaxOpenConns: dbc.SetMaxOpenConns,
		SetMaxIdleConns: dbc.SetMaxIdleConns,
		Prefix:          dbc.Prefix,
	}
	engin, err := gorose.Open(&config)
	if err != nil {
		panic(err.Error())
	}
	engin.Use(gorose.DefaultLogger())
	return engin
}
