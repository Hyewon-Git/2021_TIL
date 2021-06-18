package service

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type KubeService struct {
	provider  string
	clientSet kubernetes.Interface
	config    *rest.Config
}

func (ks *KubeService) GetClientSet(provider string) (kubernetes.Interface, error) {
	var err error
	providers := map[string]string{
		"GKE": "GKE",
		"EKS": "EKS",
	}
	// wrong provider access
	if providers[provider] == "" {
		errMsg := fmt.Sprintf("%s provider not found. Please correct provider \n prvoders = [GKE ,EKS]", provider)
		log.Print("Wrong provider")
		return nil, errors.New(errMsg)
	}
	if ks.clientSet == nil {
		ks.getClusterConfig(provider)
		ks.clientSet, err = kubernetes.NewForConfig(ks.config)
		Checkerror(err)
	}
	return ks.clientSet, nil
}

func (ks *KubeService) getClusterConfig(provider string) {
	ks.provider = provider
	var err error
	if provider == "GKE" {
		// return gke kubeconfig
	}
	if provider == "EKS" {
		// return eks kubeconfig
	}
	//일단 무조건 local config로!
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	ks.config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	Checkerror(err)
}
