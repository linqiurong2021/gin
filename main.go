package main

import (
	"html/template"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

// LoginForm 登录表单
type LoginForm struct {
	Username string `form:"username"` // 结构体字段需要大写  与表单字段绑定 from:"username"
	Password string `form:"password"`
}

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
	// 模板函数 函数需要写在解析模板前
	r.GET("/user/user_func", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user_func.tmpl", gin.H{
			"title": `<a href='http://www.baidu.com' target='_blank'>百度</a>`,
		})
	})
	//
	r.GET("/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Posts website",
		})
	})
	// 测试模板
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
	// Query 参数获取
	r.GET("/query", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "18")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	// 表单提交 GET
	r.GET("/post_form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post_form.tmpl", nil)
	})
	// 表单提交 POST
	r.POST("/post_form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": username + "----" + password,
		})
	})
	// URI
	r.GET("/uri/:username/:age", func(c *gin.Context) {
		username := c.Param("username")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
		})
	})
	// 参数绑定
	r.GET("/should_bind", func(c *gin.Context) {
		var loginForm LoginForm
		err := c.ShouldBind(&loginForm) //支持post get
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"username": loginForm.Username,
				"password": loginForm.Password,
			})
		}

	})
	// 文件上传 GET
	r.GET("/file_upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "file_upload.tmpl", nil)
	})
	// 文件上传 POST
	r.POST("/file_upload", func(c *gin.Context) {
		f, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			dir := "./uploads/files/"
			// 判断上传目录是否存在 如果不存在则创建文件夹
			_, err := os.Stat(dir)
			if os.IsNotExist(err) {
				// 递归创建  os.Mkdir 创建单个
				os.MkdirAll(dir, os.ModePerm)
			}

			distDir := path.Join(dir, f.Filename)
			c.SaveUploadedFile(f, distDir)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	// 多文件上传 GET
	r.GET("/file_muti", func(c *gin.Context) {
		c.HTML(http.StatusOK, "file_muti.tmpl", nil)
	})
	// 多文件上传 POST
	r.POST("/file_muti", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]
		dir := "./uploads/files/"
		// 判断上传目录是否存在 如果不存在则创建文件夹
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			// 递归创建  os.Mkdir 创建单个
			os.MkdirAll(dir, os.ModePerm)
		}
		for _, file := range files {
			distDir := path.Join(dir, file.Filename)
			c.SaveUploadedFile(file, distDir)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
