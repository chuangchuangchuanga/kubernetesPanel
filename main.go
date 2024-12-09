package main

import (
	"github.com/gin-gonic/gin"
	ownInformers "kubernetesPanel/informers"
	routes "kubernetesPanel/route"
)

type SharedService struct {
	Data string
}

func main() {
	ownInformers.InitInformerManager()

	r := gin.Default()
	apiGroup := r.Group("/api")
	routes.KubernetsRoute(apiGroup)
	r.Run(":8080")
}
