package ownInformers

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"sync"
	"time"
)

type informerManager struct {
	clientSet         *kubernetes.Clientset
	informerFactory   *informers.SharedInformerFactory
	podInformer       *cache.SharedIndexInformer
	namespaceInformer *cache.SharedIndexInformer
}

var (
	once     sync.Once
	instance *informerManager
)

func GetInformer() *informerManager {
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
		podInformer := informerFactory.Core().V1().Pods().Informer()
		namespaceInformer := informerFactory.Core().V1().Namespaces().Informer()

		instance = &informerManager{
			clientSet:         clientSet,
			informerFactory:   &informerFactory,
			podInformer:       &podInformer,
			namespaceInformer: &namespaceInformer,
		}
	})
	return instance
}

func (i *informerManager) Run() {

}
