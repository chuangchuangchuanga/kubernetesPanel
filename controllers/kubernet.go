package controllers

import (
	"bufio"
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

	deployment, err := ownInformers.GetInformer().GetClientSet().AppsV1().Deployments(namespaceName).Get(context.TODO(), deploymentName, metaV1.GetOptions{})
	if err != nil {
		panic(err)
	}
	labelSelector := ""
	for key, value := range deployment.Spec.Template.GetLabels() {
		if labelSelector != "" {
			labelSelector += ","
		}
		labelSelector += fmt.Sprintf("%s=%s", key, value)
	}

	deploymentPodList, err := ownInformers.GetInformer().GetClientSet().CoreV1().Pods(namespaceName).List(context.TODO(), metaV1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(200, utils.StandardResponse{}.Fail(400, err.Error(), nil))
		return
	}
	deploymentPodListVoRes := []Response.DeploymentPodVoRes{}

	for _, i := range deploymentPodList.Items {
		deploymentPodListVoRes = append(deploymentPodListVoRes, Response.DeploymentPodVoRes{
			Name:       i.Name,
			Status:     string(i.Status.Phase),
			CreateTime: i.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})

	}

	c.JSON(200, utils.StandardResponse{}.Success(deploymentPodListVoRes))
	return
}

// GetPodLogsHandler 处理 Pod 日志的 WebSocket 请求
// 该函数通过 WebSocket 提供 Pod 的实时日志流
func GetPodLogsHandler(c *gin.Context) {
	// 创建一个允许所有来源的 WebSocket Upgrader
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all connections (you can customize this)
		},
	}

	// 将 HTTP 连接升级为 WebSocket 连接
	connect, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(400, utils.StandardResponse{}.Fail(400, "Upgrade failed", nil))
		return
	}
	defer connect.Close()

	// 创建一个广播通道，用于发送日志消息
	var broadcast = make(chan string, 1000)

	// 读取客户端发送的第一个消息，该消息包含 Pod 的命名空间和名称
	_, firstMessage, err := connect.ReadMessage()
	if err != nil {
		log.Println("read first message:", err)
		return
	}

	// 解析第一个消息，获取 Pod 的命名空间和名称
	var params map[string]string
	if err := json.Unmarshal(firstMessage, &params); err != nil {
		log.Println("unmarshal first message:", err)
		return
	}

	namespace := params["namespace"]
	podname := params["podname"]

	// 启动一个 goroutine，用于将广播通道中的日志消息写入 WebSocket 连接
	go func() {
		for {
			message := <-broadcast
			err := connect.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}()

	// 获取 Pod 的日志流
	tailLines := int64(2000)
	podLogs, err := ownInformers.GetInformer().GetClientSet().CoreV1().Pods(namespace).GetLogs(podname, &v1.PodLogOptions{
		TailLines: &tailLines,
		Follow:    true,
	}).Stream(context.TODO())
	if err != nil {
		log.Fatalf("Error in log request: %v", err)
	}
	defer podLogs.Close()

	// 启动一个 goroutine，用于读取 Pod 日志流的每一行，并将其发送到广播通道
	go func() {
		scanner := bufio.NewScanner(podLogs) // 使用 Scanner 逐行读取
		for scanner.Scan() {
			if scanner.Scan() {
				// 将每一行日志发送到 broadcast 通道
				broadcast <- scanner.Text()
			} else {
				// 如果没有更多数据，检查是否到达流末尾
				if err := scanner.Err(); err != nil {
					fmt.Println("Error reading logs:", err)
				}
				break // 结束循环
			}
		}
	}()

	// 保持函数运行，以维持 WebSocket 连接
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
