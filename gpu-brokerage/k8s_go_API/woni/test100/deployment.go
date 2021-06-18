package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	typeappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/util/retry"
)

func setDeployment(clientset *kubernetes.Clientset) (deploymentsClient typeappsv1.DeploymentInterface) {
	deploymentsClient = clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	return
}

// #1. deployment 조회
func listDeployment(clientset *kubernetes.Clientset) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	Checkerror(err)

	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}
}

// #2. deployment 생성
func createDeployment(clientset *kubernetes.Clientset, addDeploymentName string) {
	// deployment 구성하기 - 매니페스트 설정

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	///삭제를요함!!!!!!!!!!!!!!!!!!!!1-------------------------------------------
	//새로 생성하는 것이므로 매니페스트 설정이 필요하다 !! deploymentsClient typev1.DeploymentInterface, >>  면한꺼번에 묶으려
	dpManifest := infoDeployment(addDeploymentName)

	//생성 Deployment
	fmt.Println("Creating Deployment !!")
	result, err := deploymentsClient.Create(context.TODO(), dpManifest, metav1.CreateOptions{})
	Checkerror(err)
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

// #3. deployment 변경
func updateDeployment(clientset *kubernetes.Clientset) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := deploymentsClient.Get(context.TODO(), "demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}

		result.Spec.Replicas = int32Ptr(1)                           // reduce replica count
		result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
		_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	fmt.Println("Updated deployment...")
}

// #4. deployment 삭제
func deleteDeployment(clientset *kubernetes.Clientset, deleteDeploymentName string) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete(context.TODO(), deleteDeploymentName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted deployment : ", deleteDeploymentName)
}

//매니페스트 정보생성
func infoDeployment(deploymentName string) (deployment *appsv1.Deployment) {
	deployment = &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2), // 생성하고 유지할 파드 개수
			Selector: &metav1.LabelSelector{ // # 디플로이먼트의 스펙에 해당!!
				MatchLabels: map[string]string{ // matchLabels :컨트롤러와 파드를 대응시키는 라벨
					"app": "demo", //파드에 해당 라벨이 있어야한다!!
				},
			},
			Template: apiv1.PodTemplateSpec{ //파드에 대한 스팩!
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo", // =파드의라벨 ( 컨트롤러의 matchLabels와 일치해야함!)
					},
				},
				Spec: apiv1.PodSpec{ //컨테이너의펙 스
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	return
}

func int32Ptr(i int32) *int32 { return &i }

// # 여러기능을 붙여서 쓸때 "Press Return key " 를해야 다음 단계로 넘어간다!!
func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}
