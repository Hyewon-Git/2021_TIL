package controller

import (
	"fmt"

	"example.com/soonbee/apiserver/service"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/gommon/log"
)

type JobController struct{}

var jobService service.JobService
var kubeService service.KubeService

func (c JobController) Init(g *echo.Group) {
	kubeService = service.KubeService{}
	jobService = service.JobService{}

	g.GET("/:namespace/job", c.GetList)
	g.GET("/:namespace/job/:name", c.Get)
	g.PUT("/:namespace/job", c.Create)
	g.POST("/:namespace/job/:name", c.Update)
	g.DELETE("/:namespace/job/:name", c.Delete)
	// ?provider=GKE or EKS
}

func (JobController) GetList(c echo.Context) error {
	clientset, err := kubeService.GetClientSet(c.QueryParam("provider"))
	if err != nil {
		//provider error
		return c.String(404, err.Error())
	}

	statusCode, jobList := jobService.GetList(clientset, c.Param("namespace"))
	return c.String(statusCode, fmt.Sprintf("Get jobList : %d , %s", statusCode, jobList))
}

func (JobController) Get(c echo.Context) error {
	clientset, err := kubeService.GetClientSet(c.QueryParam("provider"))
	if err != nil {
		return c.String(404, err.Error())
	}

	statusCode, jobpodList := jobService.Get(clientset, c.Param("namespace"), c.Param("name"))
	return c.String(statusCode, fmt.Sprintf("Get specific job - pods : %d , %s", statusCode, jobpodList))
}

func (JobController) Create(c echo.Context) error {
	clientset, err := kubeService.GetClientSet(c.QueryParam("provider"))
	if err != nil {
		return c.String(404, err.Error())
	}

	var jobContent service.Job
	if err = c.Bind(&jobContent); err != nil {
		return err
	}
	statusCode, result := jobService.Create(clientset, c.Param(("namespace")), jobContent)
	return c.String(statusCode, fmt.Sprintf("Create job : \n %d , %s", statusCode, result))
}

func (JobController) Update(c echo.Context) error {
	clientset, err := kubeService.GetClientSet(c.QueryParam("provider"))
	if err != nil {
		return c.String(404, err.Error())
	}

	var jobContent service.Job
	if err = c.Bind(&jobContent); err != nil {
		return err
	}

	stautsCode, result := jobService.Update(clientset, c.Param("namespace"), c.Param("name"), jobContent)
	return c.String(stautsCode, fmt.Sprintf("Updage job : %d, %s", stautsCode, result))
}

func (JobController) Delete(c echo.Context) error {
	clientset, err := kubeService.GetClientSet(c.QueryParam("provider"))
	if err != nil {
		return c.String(404, err.Error())
	}

	statusCode, result := jobService.Delete(clientset, c.Param("namespace"), c.Param("name"))
	return c.String(statusCode, fmt.Sprintf("Delete job : %d , %s", statusCode, result))
}
