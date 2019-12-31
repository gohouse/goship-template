package booter

import (
	"github.com/gohouse/gorose/v2"
)

type AppConfig map[string]interface{}

func (AppConfig) TableName() string {
	return "config"
}


func NewConfig() *AppConfig {
	return &AppConfig{}
}

func (c *AppConfig) GetConfig(key string) interface{} {
	if res, ok := (*c)[key]; ok {
		return res
	}
	return nil
}

func BootConfig(db *gorose.Engin) *AppConfig {
	var conf = NewConfig()

	// 加载数据库配置
	loadConfigFromDb(conf,db)

	//// 加载文件配置
	//loadConfigFromFile(conf)

	return conf
}

func loadConfigFromDb(conf *AppConfig,db *gorose.Engin) {
	//v, err := db.NewOrm().Table("config").Pluck("value", "module")
	//if err != nil {
	//	panic(err.Error())
	//}
	//for k2,v2:= range v.(map[interface{}]interface{}){
	//	(*conf)[k2.(string)] = v2
	//}
}
