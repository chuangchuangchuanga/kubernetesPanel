package main

import (
	"github.com/gin-gonic/gin"
	ownInformers "kubernetesPanel/informers"
	routes "kubernetesPanel/route"
)

func main() {
	go ownInformers.GetInformer()

	r := gin.Default()
	apiGroup := r.Group("/api")
	routes.KubernetsRoute(apiGroup)
	r.Run(":8080")
}
