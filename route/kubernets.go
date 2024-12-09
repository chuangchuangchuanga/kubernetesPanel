package routes

import "github.com/gin-gonic/gin"

func KubernetsRoute(router *gin.RouterGroup) {
	router.GET("/namespace", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World111",
		})
	})
}
