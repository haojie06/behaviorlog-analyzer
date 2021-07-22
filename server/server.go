package server

import (
	"behaviorlog-analyzer/utils"

	"github.com/gin-gonic/gin"
)

func Start(ip string, port string) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(ip + ":" + port)
	utils.CheckErr(err, "启动web服务器")
}
