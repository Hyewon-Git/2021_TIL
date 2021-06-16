package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	gobatchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type JobService struct{}

type Job struct {
	Name        string   `json:"name"`
	Namespace   string   `json:"namespace"`
	Gputype     string   `json:"gputype"`
	Tags        []string `json:"tags`
	Description string   `json:"description`
	//container info
	Image   string   `json:"image"`
	Command string   `json:"command"`
	Args    []string `json:"args"`

	CPU      float32 `json:"cpu"`
	RAM      float32 `json:"ram"`
	GpuCount int     `json:"gpuCount"`
}

func (js JobService) GetList(clientset kubernetes.Interface, namespace string) (int, string) {
	// return job list
	jobClient := js.getJobClient(clientset, namespace)
	jobList, err := jobClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return ReturnError(err)
	}
	log.Print("Get job list ")
	return 200, GetJSONStr(jobList)
}

func (js JobService) Get(clientset kubernetes.Interface, namespace string, name string) (int, string) {
	// return job info & job-pods info
	jobClient := js.getJobClient(clientset, namespace)
	job, err := jobClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return ReturnError(err)
	}

	podClient := js.getJobClient(clientset, namespace)
	podList, err := podClient.List(context.TODO(), metav1.ListOptions{
		LabelSelector: "job-name=" + name,
	})
	if err != nil {
		return ReturnError(err)
	}
	log.Print("Get specific job & pods ")
	return 200, fmt.Sprint("{job : ", GetJSONStr(job), " ,\n pod : ", GetJSONStr(podList.Items))
}

//dudaji.com/gputype: t4 / k80 / v100
func (js JobService) Create(clientset kubernetes.Interface, namespace string, jobContent Job) (int, string) {
	// create job
	jobClient := js.getJobClient(clientset, namespace)
	jobBody := js.jobBodyParser(jobContent)
	createJob, err := jobClient.Create(context.TODO(), jobBody, metav1.CreateOptions{})
	if err != nil {
		return ReturnError(err)
	}
	log.Print("Create job ")
	return 200, GetJSONStr(createJob)
}

func (js JobService) Update(clientset kubernetes.Interface, namespace string, name string, jobContent Job) (int, string) {
	jobClient := js.getJobClient(clientset, namespace)
	job, err := jobClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return ReturnError(err)
	}
	// Renew job info - annotations
	if jobContent.Tags != nil {
		job.ObjectMeta.Annotations["tags"] = strings.Join(jobContent.Tags, ",")
	}
	if jobContent.Description != "" {
		job.ObjectMeta.Annotations["description"] = jobContent.Description
	}
	updateJob, err := jobClient.Update(context.TODO(), job, metav1.UpdateOptions{})
	if err != nil {
		return ReturnError(err)
	}
	log.Print("Update job")
	return 200, GetJSONStr(updateJob)
}

func (js JobService) Delete(clientset kubernetes.Interface, namespace string, name string) (int, string) {
	// delete job
	jobClient := js.getJobClient(clientset, namespace)

	deletePolicy := metav1.DeletePropagationForeground
	err := jobClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		return ReturnError(err)
	}
	log.Print("Delete job")
	return 200, name
}

func (JobService) getJobClient(clientset kubernetes.Interface, namespace string) gobatchv1.JobInterface {
	return clientset.BatchV1().Jobs(namespace)
}

func (JobService) getPodClient(clientset kubernetes.Interface, namespace string) v1.PodInterface {
	return clientset.CoreV1().Pods(namespace)
}

func (JobService) jobBodyParser(jobContent Job) (newJob *batchv1.Job) {
	newJob = &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobContent.Name,
			Namespace: jobContent.Namespace,
			Labels: map[string]string{
				"gputype": jobContent.Gputype,
			},
			Annotations: map[string]string{
				"tags":        strings.Join(jobContent.Tags, ","),
				"description": jobContent.Description,
			},
		},
		Spec: batchv1.JobSpec{
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  jobContent.Name,
							Image: jobContent.Image,
							Command: []string{
								jobContent.Command,
							},
							Args: jobContent.Args,
							Resources: apiv1.ResourceRequirements{
								Limits: apiv1.ResourceList{
									"cpu":            resource.MustParse(fmt.Sprintf("%f", jobContent.CPU)),
									"memory":         resource.MustParse(fmt.Sprintf("%f", jobContent.RAM) + "Gi"),
									"nvidia.com/gpu": resource.MustParse(strconv.Itoa(jobContent.GpuCount)),
								},
							},
						},
					},
					RestartPolicy: apiv1.RestartPolicyNever,
				},
			},
		},
	}
	return
}
