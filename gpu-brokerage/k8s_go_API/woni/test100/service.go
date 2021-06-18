package main

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func setService(clientset *kubernetes.Clientset) (servicesClient typedcorev1.ServiceInterface) {
	servicesClient = clientset.CoreV1().Services(apiv1.NamespaceDefault)
	return
}

// #1. service  조회
func listService(clientset *kubernetes.Clientset) {
	servicesClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)

	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	list, err := servicesClient.List(context.TODO(), metav1.ListOptions{})
	Checkerror(err)

	for _, s := range list.Items {
		fmt.Printf(" * %s \n", s.Name)
	}
}

// #2. service  생성
func createService(clientset *kubernetes.Clientset, serviceName string, podName string) {
	// deployment 구성하기 - 매니페스트 설정

	servicesClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	svManifest := infoService(serviceName, podName)

	fmt.Println("Creating Service !!")
	result, err := servicesClient.Create(context.TODO(), svManifest, metav1.CreateOptions{})
	Checkerror(err)
	fmt.Printf("Created Service %q.\n", result.GetObjectMeta().GetName())
}

// #3. service 뱐경

// #4. service 삭제

func infoService(serviceName string, podName string) (swManifest *apiv1.Service) {
	swManifest = &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName, // 서비스 이름
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{ // 서비스 - 파드 매핑하는 설정
				"app": podName, //파드의 이름이여야함!

			},
			Type: "NodePort",
			Ports: []apiv1.ServicePort{ //서비스 공개방법 = 생략하면 ClusterIP  사용
				{
					Protocol: apiv1.ProtocolTCP,
					Port:     80,
					NodePort: 30001,
				},
			},
		},
	}
	return
}
