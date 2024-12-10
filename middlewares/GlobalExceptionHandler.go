package middlewares

import (
	"github.com/gin-gonic/gin"
	"kubernetesPanel/utils"
	"log"
)

func GlobalExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic occurred: %v", err)
				c.JSON(
					200, utils.StandardResponse{}.Fail(500, "Server Error", nil),
				)
				c.Abort()
			}
		}()
		c.Next()
	}
}
