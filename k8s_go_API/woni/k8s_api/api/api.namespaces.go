package api

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func LsNamespace(clientset *kubernetes.Clientset, namespace string) bool {
	nsList := ListNamespace(clientset)
	for _, ns := range nsList {
		if *&ns.Name == namespace {
			return true
		}
	}
	return false
}

func ListNamespace(clientset *kubernetes.Clientset) []v1.Namespace {
	nsList, _ := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	for _, ns := range nsList.Items {
		fmt.Printf(" * %s \n", *&ns.Name)
	}
	return nsList.Items
}
func CreateNamespace(clientset *kubernetes.Clientset, namespace string) {
	fmt.Println("createnamespace")
	namespacesClient := clientset.CoreV1().Namespaces()
	nsManifest := InfoNamespace(namespace)

	//생성 Deployment
	result, err := namespacesClient.Create(context.TODO(), nsManifest, metav1.CreateOptions{})
	Checkerror(err)
	fmt.Printf("Created Namespace %q.\n", result.Name)

}

func InfoNamespace(namespace string) (nsManifest *v1.Namespace) {
	nsManifest = &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace, // 서비스 이름
		},
	}
	return

}
