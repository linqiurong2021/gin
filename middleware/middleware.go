package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthCheck 权限校验中间件
func AuthCheck(check bool) gin.HandlerFunc {
	// 查询数据库等
	fmt.Println("middleware")
	return func(c *gin.Context) {
		//
		_, ok := c.Get("token")
		if !ok && check {
			fmt.Println("middleware check failure")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token check faliure",
			})
			c.Abort() // 阻止后面的函数
		} else {
			fmt.Println("middleware check success")
			c.Next() // 调用后面的函数
		}
	}
}
