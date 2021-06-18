package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func Checkerror(err error) {
	if err != nil {
		fmt.Println("Checkerror here")
		panic(err)
	}
}
func main() {
	// kubeconfig에서 현재 콘텍스트를 사용한다
	// 현재폴더 주소 가져오기
	var kubeconfig *string
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

	// clientset을 생성한다
	clientset, err := kubernetes.NewForConfig(config)
	Checkerror(err)

	// #1. 현재 실행중인 파드 조회
	// 파드를 나열하기 위해 API에 접근한다
	pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	// #1. 현재 실행중은 deployment 조회
	listDeployment(clientset)
	// #2. deployment 생성
	createDeployment(clientset, "hyewon-dp-02")
	listDeployment(clientset)
	// #3. deployment 삭제
	deleteDeployment(clientset, "hyewon-dp-01")

	createService(clientset, "hyewon-sv-01", "demo")
	listService(clientset)

}
