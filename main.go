package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	ownInformers "kubernetesPanel/informers"
	"kubernetesPanel/middlewares"
	routes "kubernetesPanel/route"
	"log"
	"net/http"
	"os"
)

type SharedService struct {
	Data string
}

var f embed.FS

func main() {
	ownInformers.InitInformerManager()

	r := gin.Default()
	r.Use(middlewares.GlobalExceptionHandler())
	r.Use(middlewares.Serve("/", middlewares.EmbedFolder(f, "web/assets")))
	r.NoRoute(func(c *gin.Context) {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Current working directory:", currentDir)
		data, err := f.ReadFile("web/index.html")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})
	apiGroup := r.Group("/api")
	routes.KubernetsRoute(apiGroup)
	r.Run(":8080")
}
