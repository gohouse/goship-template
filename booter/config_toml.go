package booter

import (
	"github.com/BurntSushi/toml"
	"github.com/gohouse/gosms/adapter/sdks"
	"io/ioutil"
)

type BaseConfig struct {
	SiteInfo      SiteInfo          `toml:"site_info" json:"site_info"`
	Jwt           Jwt               `toml:"jwt" json:"jwt"`
	Error         Error             `toml:"error" json:"error"`
	Database      Database          `toml:"database" json:"database"`
	CasbinModel   map[string]string `toml:"casbin_model" json:"casbin_model"`
	AliyunOptions *sdks.AliyunOptions
	TwilioOptions *sdks.TwilioOptions
	AliyunOSS     AliyunOSS
	Host          Host
}
type Host struct {
	Chapter string
	Img     string
}
type AliyunOSS struct {
	AccessKeyId   string
	AccessSecret  string
	Endpoint      string
	DefaultBucket string
}
type SiteInfo struct {
	Env        string `toml:"env" json:"env"`
	HttpPort   string `toml:"http_port" json:"http_port"`
	UploadPath string `toml:"upload_path" json:"upload_path"`
}
type Jwt struct {
	Secret string `toml:"secret" json:"secret"`
	Expire int64  `toml:"expire" json:"expire"`
}
type Error struct {
	ErrorStackLayer int `toml:"error_stack_layer" json:"error_stack_layer"`
	ErrorLogFile string `toml:"error_log_file" json:"error_log_file"`
}
type Database struct {
	Driver string `toml:"driver" json:"driver"` // 驱动: mysql/sqlite3/oracle/mssql/postgres/clickhouse, 如果集群配置了驱动, 这里可以省略
	// mysql 示例:
	// root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=true
	Dsn             string `toml:"dsn" json:"dsn"`                         // 数据库链接
	SetMaxOpenConns int    `toml:"setMaxOpenConns" json:"setMaxOpenConns"` // (连接池)最大打开的连接数，默认值为0表示不限制
	SetMaxIdleConns int    `toml:"setMaxIdleConns" json:"setMaxIdleConns"` // (连接池)闲置的连接数, 默认0
	Prefix          string `toml:"prefix" json:"prefix"`                   // 表前缀, 如果集群配置了前缀, 这里可以省略
}

func LoadBaseConfig(file string) *BaseConfig {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	var config BaseConfig
	err = toml.Unmarshal(bytes, &config)
	if err != nil {
		panic(err.Error())
	}
	return &config
}
