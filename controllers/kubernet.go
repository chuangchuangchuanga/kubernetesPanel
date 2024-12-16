package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubernetesPanel/vo/Request"
	"kubernetesPanel/vo/Response"
	"log"
	"net/http"
	"time"

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
	return
}

func GetDeploymentHandler(c *gin.Context) {
	var req Request.DeploynetListVoReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, utils.StandardResponse{}.Fail(400, err.Error(), nil))
		return
	}

	namespace := req.GetNamespace()
	deploymentList, err := ownInformers.GetInformer().GetClientSet().AppsV1().Deployments(namespace).List(context.TODO(), metaV1.ListOptions{})
	deploymentitems := deploymentList.Items
	var deploymentVoRes []Response.DeploymentVoRes
	for _, i := range deploymentitems {
		deploymentVoRes = append(deploymentVoRes, Response.DeploymentVoRes{Name: i.Name, Namespace: namespace})
	}
	if err != nil {
		panic(err)
	}

	c.JSON(200, utils.StandardResponse{}.Success(deploymentVoRes))
	return
}

func GetDeployemntPodHandler(c *gin.Context) {
	var req Request.DeploynetPodListVoReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, utils.StandardResponse{}.Fail(400, err.Error(), nil))
		return
	}

	deploymentName := req.GetDeploymentName()
	namespaceName := req.GetNamespace()

	deploymentPodList, err := ownInformers.GetInformer().GetClientSet().CoreV1().Pods(namespaceName).List(context.TODO(), metaV1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", deploymentName),
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(200, utils.StandardResponse{}.Fail(400, err.Error(), nil))
		return
	}
	deploymentPodListVoRes := Response.DeploymentPodVoRes{}

	for _, i := range deploymentPodList.Items {
		deploymentPodListVoRes.AddName(i.Name)
	}

	c.JSON(200, utils.StandardResponse{}.Success(deploymentPodListVoRes.GetName()))
	return
}

func GetPodLogsHandler(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all connections (you can customize this)
		},
	}

	connect, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(400, utils.StandardResponse{}.Fail(400, "Upgrade failed", nil))
		return
	}
	defer connect.Close()
	var broadcast = make(chan string, 1000)

	_, firstMessage, err := connect.ReadMessage()
	if err != nil {
		log.Println("read first message:", err)
		return
	}

	var params map[string]string
	if err := json.Unmarshal(firstMessage, &params); err != nil {
		log.Println("unmarshal first message:", err)
		return
	}

	namespace := params["namespace"]
	podname := params["podname"]

	go func() {
		i := 0
		for {
			i++
			message := <-broadcast
			err := connect.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}()

	podLogs, err := ownInformers.GetInformer().GetClientSet().CoreV1().Pods(namespace).GetLogs(podname, &v1.PodLogOptions{
		Follow: true,
	}).Stream(context.TODO())
	if err != nil {
		log.Fatalf("Error in log request: %v", err)
	}
	defer podLogs.Close()
	go func() {

		buf := make([]byte, 2024)
		for true {
			n, err := podLogs.Read(buf)
			if err != nil {
				fmt.Println("Error reading logs:", err)
			}
			broadcast <- string(buf[:n])
		}
	}()
	for true {
		time.Sleep(1000 * time.Second)
	}
}

func RestartDeploymentHandler(c *gin.Context) {
	var req Request.DeploymentRestartVoReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, utils.StandardResponse{}.Fail(400, err.Error(), nil))
		return
	}

	namespace := req.GetNamespace()
	deploymentName := req.GetDeploymentName()
	deploymentClient := ownInformers.GetInformer().GetClientSet().AppsV1().Deployments(namespace)
	deployment, err := deploymentClient.Get(context.TODO(), deploymentName, metaV1.GetOptions{})

	if err != nil {
		c.JSON(200, utils.StandardResponse{}.Fail(400, err.Error(), nil))
		return
	}
	if deployment.Spec.Template.Annotations == nil {
		deployment.Spec.Template.Annotations = make(map[string]string)
	}
	deployment.Spec.Template.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)

	_, err = deploymentClient.Update(context.TODO(), deployment, metaV1.UpdateOptions{})
	if err != nil {
		c.JSON(200, utils.StandardResponse{}.Fail(400, err.Error(), nil))
		return
	}

	c.JSON(200, utils.StandardResponse{}.Success(nil))
	return

}
