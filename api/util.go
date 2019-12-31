package api

import (
	"github.com/gohouse/gorose/v2"
	"goshipdemo/util"
)

func DB() gorose.IOrm {
	return util.DB()
}