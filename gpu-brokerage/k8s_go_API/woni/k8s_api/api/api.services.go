package api

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

//나중에 여러기능 한꺼번에 사용할때 먼저 선언해서  Client초기화 시키면됨!
func SetService(clientset *kubernetes.Clientset) (servicesClient typedcorev1.ServiceInterface) {
	servicesClient = clientset.CoreV1().Services(apiv1.NamespaceDefault)
	return
}

// #1. service  조회
func ListService(clientset *kubernetes.Clientset, namespace string) []apiv1.Service {
	servicesClient := clientset.CoreV1().Services(namespace)

	fmt.Printf("Listing services in namespace %q:\n", namespace)
	list, err := servicesClient.List(context.TODO(), metav1.ListOptions{})
	Checkerror(err)

	for _, s := range list.Items {
		fmt.Printf(" * %s \n", s.Name)
	}
	return list.Items
}

// #2. service  생성
func CreateService(clientset *kubernetes.Clientset, sv ServiceStruct) {
	// deployment 구성하기 - 매니페스트 설정
	servicesClient := clientset.CoreV1().Services(sv.SNS)
	svManifest := InfoService(sv)

	if !LsNamespace(clientset, sv.SNS) {
		CreateNamespace(clientset, sv.SNS)
	}
	result, err := servicesClient.Create(context.TODO(), svManifest, metav1.CreateOptions{})
	Checkerror(err)
	fmt.Printf("Created Service %q in namespace : %q", result.GetObjectMeta().GetName(), result.GetObjectMeta().GetNamespace())
}

// #3. service 뱐경
func DeleteService(clientset *kubernetes.Clientset, sv ServiceStruct) {
	servicesClient := clientset.CoreV1().Services(sv.SNS)

	deletePolicy := metav1.DeletePropagationForeground
	if err := servicesClient.Delete(context.TODO(), sv.SNAME, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted service : ", sv.SNAME)
}

// #4. service 삭제
func InfoService(sv ServiceStruct) (swManifest *apiv1.Service) {
	swManifest = &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sv.SNAME, // 서비스 이름
			Namespace: sv.SNS,
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{ // 서비스 - 파드 매핑하는 설정
				"app": "demo", //파드의 이름이여야함!

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

type ServiceStruct struct {
	SNAME string `json:"sname"`
	SNS   string `json:"sns`
	SPOD  string `json:"spod`
}
