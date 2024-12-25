package ownInformers

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type InformerManager struct {
	clientSet                *kubernetes.Clientset
	informerFactory          informers.SharedInformerFactory
	podInformerStore         cache.Store
	namespaceInformerStore   cache.Store
	deploymentsInformerStore cache.Store
}

func InitInformerManager() {
	informerManager := GetInformer()
	go informerManager.Run(context.Background())
}

var (
	once     sync.Once
	instance *InformerManager
)

func GetInformer() *InformerManager {
	once.Do(func() {
		var kubeconfig *string
		if home := homeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		// 尝试使用集群内配置
		config, err := rest.InClusterConfig()
		if err != nil {
			// 集群内配置失败，尝试使用本地 kubeconfig 文件
			config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
			if err != nil {
				fmt.Printf("Error building kubeconfig: %s\n", err.Error())
				os.Exit(1)
			}
		}

		if err != nil {
			panic(err)
		}
		clientSet, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}

		informerFactory := informers.NewSharedInformerFactoryWithOptions(clientSet, 100*time.Second)
		podInformerStore := informerFactory.Core().V1().Pods().Informer().GetStore()
		deploymentsInformerStore := informerFactory.Apps().V1().Deployments().Informer().GetStore()
		namespaceInformerStore := informerFactory.Core().V1().Namespaces().Informer().GetStore()

		instance = &InformerManager{
			clientSet:                clientSet,
			informerFactory:          informerFactory,
			podInformerStore:         podInformerStore,
			namespaceInformerStore:   namespaceInformerStore,
			deploymentsInformerStore: deploymentsInformerStore,
		}
	})
	return instance
}

func (i *InformerManager) Run(ctx context.Context) {
	i.informerFactory.Start(ctx.Done())
	i.informerFactory.WaitForCacheSync(ctx.Done())
}

func (i *InformerManager) GetPodInformerStore() cache.Store {
	return i.podInformerStore
}

func (i *InformerManager) GetDeploymentInformerStore() cache.Store {
	return i.deploymentsInformerStore
}

func (i *InformerManager) GetClientSet() *kubernetes.Clientset {
	return i.clientSet
}

// 获取用户主目录路径
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
