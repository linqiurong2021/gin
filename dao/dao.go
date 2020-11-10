package dao

import (
	"fmt"
	"gin/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DB 数据库
	DB *gorm.DB
)

// InitDB 初始化数据库连接
func InitDB(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB, cfg.Charset)
	fmt.Printf("init DB: %s \n", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return
}
