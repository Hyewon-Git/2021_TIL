package api

import (
	"flag"
	"log"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func Checkerror(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
func ConfirmConfig() *rest.Config {
	// kubeconfig에서 현재 콘텍스트를 사용한다
	// 현재폴더 주소 가져오기
	var kubeconfig *string
	//flag로 config를 저장하면 > 해당 url로 2번이상 선언이안된다! >>set으로 고쳐서 여러번 접근간으하게 만들어줌!
	//flagset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	//이미지 구워진건 --> set으로안해서아마 두번이상 url접속 안됨!!!
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// path-to-kubeconfig -- 예를 들어, /root/.kube/config
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//config, _ := clientcmd.BuildConfigFromFlags("", "/Users/dudaji/.kube/config")
	Checkerror(err)

	return config
}

func SetClient(config *rest.Config) *kubernetes.Clientset {
	// clientset을 생성한다
	clientset, err := kubernetes.NewForConfig(config)
	Checkerror(err)

	return clientset
}
