package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	ownInformers "kubernetesPanel/informers"
)

func KubernetsRoute(router *gin.RouterGroup) {
	router.GET("/namespace", func(c *gin.Context) {
		i := ownInformers.GetInformer().GetPodInformerStore()
		for _, a := range i.List() {
			fmt.Printf("name: %+v, status: %+v\n", a.(*v1.Pod).GetName(), a.(*v1.Pod).Status.Phase)
		}
		c.JSON(200, gin.H{
			"message": "Hello World111",
		})
	})
}
