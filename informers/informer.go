package ownInformers

import (
	"context"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"sync"
	"time"
)

type InformerManager struct {
	clientSet              *kubernetes.Clientset
	informerFactory        informers.SharedInformerFactory
	podInformerStore       cache.Store
	namespaceInformerStore cache.Store
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
		config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
		if err != nil {
			panic(err)
		}
		clientSet, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}

		informerFactory := informers.NewSharedInformerFactoryWithOptions(clientSet, 100*time.Second)
		podInformerStore := informerFactory.Core().V1().Pods().Informer().GetStore()
		namespaceInformerStore := informerFactory.Core().V1().Namespaces().Informer().GetStore()

		instance = &InformerManager{
			clientSet:              clientSet,
			informerFactory:        informerFactory,
			podInformerStore:       podInformerStore,
			namespaceInformerStore: namespaceInformerStore,
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
