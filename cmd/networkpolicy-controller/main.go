package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/elisasre/networkpolicy-controller/pkg/controller"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	log.SetOutput(os.Stdout)

	sigs := make(chan os.Signal, 1)
	stop := make(chan struct{})

	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	wg := &sync.WaitGroup{}

	// Create clientset for interacting with the kubernetes cluster
	clientset, err := newClientSet()

	if err != nil {
		panic(err.Error())
	}

	configFile := "./hack/dev.json"
	if os.Getenv("CONFIG") != "" {
		configFile = os.Getenv("CONFIG")
	}
	log.Printf("Starting application...")
	go controller.NewNetworkWatcher(clientset, configFile).Run(stop, wg)
	<-sigs
	log.Printf("Shutting down...")
	close(stop)
	wg.Wait()
}

func newClientSet() (*kubernetes.Clientset, error) {
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
