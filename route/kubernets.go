package routes

import (
	"github.com/gin-gonic/gin"
	"kubernetesPanel/controllers"
)

func KubernetsRoute(router *gin.RouterGroup) {
	router.GET("/namespace", controllers.GetNamespaceHandler)
	router.POST("/deploymentlist", controllers.GetDeploymentHandler)
}
