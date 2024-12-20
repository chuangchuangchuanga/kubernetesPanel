package main

import (
	"github.com/gin-gonic/gin"
	ownInformers "kubernetesPanel/informers"
	"kubernetesPanel/middlewares"
	routes "kubernetesPanel/route"
)

type SharedService struct {
	Data string
}

func main() {
	ownInformers.InitInformerManager()

	r := gin.Default()
	r.Use(middlewares.GlobalExceptionHandler())
	apiGroup := r.Group("/api")
	routes.KubernetsRoute(apiGroup)
	r.Run(":8080")
}
