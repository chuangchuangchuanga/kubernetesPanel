package main

import (
	"github.com/gin-gonic/gin"
	ownInformers "kubernetesPanel/informers"
	"kubernetesPanel/middlewares"
	routes "kubernetesPanel/route"
)

func main() {
	ownInformers.InitInformerManager()

	r := gin.Default()
	r.Use(middlewares.GlobalExceptionHandler())
	r.Static("/assets", "./web/assets")
	r.LoadHTMLFiles("./web/index.html")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	apiGroup := r.Group("/api")
	routes.KubernetsRoute(apiGroup)
	r.Run(":8080")
}
