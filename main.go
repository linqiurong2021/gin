package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 静态文件处理 页面中有/css会 转义为 ./static/css
	// r.Static("/static", "./static")
	r.Static("/assets", "./static/assets")
	// 设置函数 需要在解析模板前
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	r.LoadHTMLGlob("templates/**/*") // 其中 ** 代表目录
	// r.LoadHTMLGlob("templates/home/*.html") //
	// r.LoadHTMLFiles("templates/user/user_func.tmpl")
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Users website",
		})
	})
	r.GET("/user/user_func", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user_func.tmpl", gin.H{
			"title": `<a href='http://www.baidu.com' target='_blank'>百度</a>`,
		})
	})

	r.GET("/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Posts website",
		})
	})

	r.GET("/home/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/home/about.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})
	r.GET("/home/contact.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})
	r.GET("/home/services.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "services.html", nil)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
