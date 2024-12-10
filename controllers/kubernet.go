package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubernetesPanel/vo/Response"

	ownInformers "kubernetesPanel/informers"
	"kubernetesPanel/utils"
)

func GetNamespaceHandler(c *gin.Context) {
	namespace, err := ownInformers.GetInformer().GetClientSet().CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		panic(err)
	}
	list := namespace.Items
	namespaceVo := Response.NamespaceVo{}
	for _, i := range list {
		namespaceVo.AddName(i.Name)
	}

	c.JSON(200, utils.StandardResponse{}.Success(namespaceVo.GetName()))
}
