package setting

import (
	"gopkg.in/ini.v1"
)

// AppConfig App配置项
type AppConfig struct {
	Release      bool `ini:"release"`
	Port         int  `ini:"port"`
	*MySQLConfig `ini:"mysql"`
}

// MySQLConfig 数据库配置项
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Charset  string `ini:"charset"`
}

// Conf 配置
var Conf = new(AppConfig)

// Init 初始化
func Init(file string) error {
	return ini.MapTo(Conf, file)
}
