package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubernetesPanel/vo/Request"
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
	namespaceVo := Response.NamespaceVoRes{}
	for _, i := range list {
		namespaceVo.AddName(i.Name)
	}

	c.JSON(200, utils.StandardResponse{}.Success(namespaceVo.GetName()))
}

func GetDeploymentHandler(c *gin.Context) {
	var req Request.DeploynetListVoReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, utils.StandardResponse{}.Fail(400, err.Error(), nil))
	}

	namespace := req.GetNamespace()
	deploymentList, err := ownInformers.GetInformer().GetClientSet().AppsV1().Deployments(namespace).List(context.TODO(), metaV1.ListOptions{})
	deploymentitems := deploymentList.Items
	deploymentVoRes := Response.DeploymentVoRes{}
	for _, i := range deploymentitems {
		deploymentVoRes.AddName(i.Name)
	}
	if err != nil {
		panic(err)
	}

	c.JSON(200, utils.StandardResponse{}.Success(deploymentVoRes.GetName()))
}
