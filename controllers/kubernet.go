package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubernetesPanel/utils"
)

func GetNamespace(c *gin.Context) {
	fmt.Printf("1111")
	c.JSON(200, utils.Response{}.Success("1111"))

}
