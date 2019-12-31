package booter

import (
	"github.com/gohouse/e"
	"github.com/gohouse/t"
)

func BootError(errconf Error) {
	err := e.NewErrorContext()

	// 获取配置中获取错误堆栈的层数
	layer := t.New(errconf.ErrorStackLayer).Int()

	// 使用中间件做错误处理,可以持久化,也可以打印到控制台等
	if errconf.ErrorLogFile != "" {
		_ = err.Use(e.LogFile("error.log"))
	}

	// 设置设置堆栈层数,如果不设置,为了节约内存占用,默认为1层
	err.Setlayer(layer)
}
