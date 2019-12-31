package main

import (
	"flag"
	"fmt"
	"log"
	"goshipdemo/booter"
	"goshipdemo/routing"
)

var (
	f string
)

func init() {
	flag.StringVar(&f, "f", "", "配置文件")
	flag.StringVar(&f, "file", "", "配置文件")
	flag.Parse()
	if f == "" {
		f = "./config.toml"
	}
}

// @title 小说站接口文档
// @version 1.0
// @description 小说站api.
// @termsOfService http://localhost:8081

// @contact.name Kevin
// @contact.url http://localhost:8081
// @contact.email fizzday@yeah.net

// @host localhost:8081
// @BasePath /api/v1
func main() {
	// 初始化驱动
	b := booter.NewBooter(f)

	// 初始化路由
	routing.Run(b.Engine)

	// 开启服务
	log.Println(fmt.Sprintf("访问: http://localhost:%s", b.SiteInfo.HttpPort))
	err := b.Engine.Run(":" + b.SiteInfo.HttpPort)

	if err != nil {
		log.Fatal(err.Error())
	}
}
