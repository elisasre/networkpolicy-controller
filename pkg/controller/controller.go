package controller

import (
	"context"
	"sync"
	"time"

	"k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

const (
	cacheTime = 3 * 60 * time.Second
)

// Controller contains controller variables that is needed to work in class
type Controller struct {
	nsInformer cache.SharedInformer
	npInformer cache.SharedInformer
	kclient    *kubernetes.Clientset
}

// Run starts the process for listening for event changes and acting upon those changes.
func (c *Controller) Run(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	// Execute go function
	go c.nsInformer.Run(stopCh)
	go c.npInformer.Run(stopCh)

	// Wait till we receive a stop signal
	<-stopCh
}

// NewNetworkWatcher creates a new nsController
func NewNetworkWatcher(kclient *kubernetes.Clientset) *Controller {
	watcher := &Controller{}
	ctx := context.Background()
	nsInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return kclient.CoreV1().Namespaces().List(ctx, options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return kclient.CoreV1().Namespaces().Watch(ctx, options)
			},
		},
		&v1.Namespace{},
		cacheTime,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	nsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: watcher.createNS,
	})

	npInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return kclient.NetworkingV1().NetworkPolicies(v1.NamespaceAll).List(ctx, options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return kclient.NetworkingV1().NetworkPolicies(v1.NamespaceAll).Watch(ctx, options)
			},
		},
		&networkv1.NetworkPolicy{},
		cacheTime,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	npInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: watcher.deleteNP,
	})

	watcher.kclient = kclient
	watcher.nsInformer = nsInformer
	watcher.npInformer = npInformer
	return watcher
}

func (c *Controller) createNS(obj interface{}) {
	ns := obj.(*v1.Namespace)
	c.ensurePolicyExist(ns.Name)
}

func (c *Controller) deleteNP(obj interface{}) {
	np := obj.(*networkv1.NetworkPolicy)
	c.ensurePolicyExist(np.Namespace)
}
