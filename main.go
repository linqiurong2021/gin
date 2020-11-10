package main

import (
	"fmt"
	"gin/dao"
	"gin/models"
	"gin/routers"
	"gin/setting"
	"os"
)

func main() {
	// 判断是否使用配置文件参数
	if len(os.Args) < 2 {
		fmt.Println("Usage：./bubble conf/config.ini")
	}
	// 加载配置文件(这里可以使用默认的配置文件)
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 初始化数据库
	err := dao.InitDB(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 关联数据表
	dao.DB.AutoMigrate(&models.Todo{})
	//
	router := routers.Init()
	// 初始化
	port := fmt.Sprintf(":%d", setting.Conf.Port)

	if err := router.Run(port); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
