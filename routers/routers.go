package routers

import (
	"gin/controller"
	"gin/setting"

	"github.com/gin-gonic/gin"
)

// Init 初始化
func Init() (r *gin.Engine) {

	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r = gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	//
	r.GET("/", controller.Index, nil)

	v1Group := r.Group("/v1")
	{
		v1Group.GET("/todo", controller.GetList)
		v1Group.POST("/todo", controller.Create)
		v1Group.PUT("/todo/:id", controller.Update)
		v1Group.DELETE("/todo/:id", controller.Delete)
	}
	return
}
