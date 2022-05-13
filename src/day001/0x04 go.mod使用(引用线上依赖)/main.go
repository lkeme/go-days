package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建默认路由
	r := gin.Default()
	// 定义路由规则和行为  http://host:port/ping  -> {"message": "pong"}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000") // 监听并在 0.0.0.0:8000 上启动服务
}
