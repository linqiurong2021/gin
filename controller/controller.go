package controller

import (
	"fmt"
	"gin/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

//Index 首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// Create 创建新的
func Create(c *gin.Context) {
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "",
			"msg":  "params invalidate",
			"code": 400,
		})
	}
	err = models.Create(&todo)
	if err != nil {
		fmt.Printf("create todo error : %s", err)
	}
	// 成功返回数据
	c.JSON(http.StatusOK, gin.H{
		"data": todo,
		"msg":  "create success",
		"code": 200,
	})
}

// Update 更新
func Update(c *gin.Context) {
	// 获取参数并查找当前数据
	// 如果存在则更新 如果不存在
	id, ok := c.Params.Get("id")
	fmt.Printf("ID:%s\n", id)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "",
			"msg":  "params invalidate",
			"code": 400,
		})
		return
	}
	todo, err := models.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "",
			"msg":  "ID invalidate",
			"code": 400,
		})
		return
	}
	c.BindJSON(&todo)
	if err = models.Update(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "",
			"msg":  "ID invalidate",
			"code": 400,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": todo,
		"msg":  "update success",
		"code": 200,
	})
	return

}

// GetList 获取列表
func GetList(c *gin.Context) {
	var list []*models.Todo
	list, err := models.GetAllList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": list,
			"msg":  err,
			"code": 400,
		})
	}
	// 成功返回数据
	c.JSON(http.StatusOK, gin.H{
		"data": list,
		"msg":  "create success",
		"code": 200,
	})
}

// Delete 删除
func Delete(c *gin.Context) {
	id, ok := c.Params.Get("id") // :id
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "",
			"msg":  "ID invalidate",
			"code": 400,
		})
		return
	}
	if err := models.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "",
			"msg":  err,
			"code": 400,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": "",
			"msg":  "create success",
			"code": 200,
		})
	}
}
