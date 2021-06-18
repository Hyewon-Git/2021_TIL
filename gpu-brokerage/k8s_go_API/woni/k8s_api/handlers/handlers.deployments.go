package handlers

import (
	"fmt"
	"k8s_api/api"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDeployment(c *gin.Context) {
	np := c.Param("namespace")
	fmt.Println(np)
	if np == "all" {
		np = ""
		//c.JSON(http.StatusBadRequest, "Please request namespace")
		//np = apiv1.NamespaceDefault
	}
	config := api.ConfirmConfig()
	clientset := api.SetClient(config)
	result := api.ListDeployment(clientset, np)

	c.JSON(http.StatusOK, result)
}

// createDeployment (PUT으로 다음 데이터형식으로!)
// DNAME  string `json:"dname"`
// DNAMESPACE string `json:"dnamespace`
func PutDeployment(c *gin.Context) {
	var createDp api.DeploymentStruct
	//var newitem models.ItemStruct
	err := c.Bind(&createDp) //  Put으로입력받은 body값들을 newitem에 바인드시김
	if err != nil {
		log.Fatal(err)
		fmt.Println("bind error!")
	}

	config := api.ConfirmConfig()
	clientset := api.SetClient(config)
	api.CreateDeployment(clientset, createDp)

	//fmt.Printf("Created deployment %q.\n", dpName)
	c.JSON(http.StatusOK, createDp)

}

func DeleteDeployment(c *gin.Context) {
	var deleteDp api.DeploymentStruct
	//var newitem models.ItemStruct
	err := c.Bind(&deleteDp) //  Put으로입력받은 body값들을 newitem에 바인드시김
	if err != nil {
		log.Fatal(err)
		fmt.Println("bind error!")
	}
	config := api.ConfirmConfig()
	clientset := api.SetClient(config)
	api.DeleteDeployment(clientset, deleteDp)
	result := fmt.Sprintf("namespace : %s \ndeployment : %s \n delete", deleteDp.DNS, deleteDp.DNAME)
	c.JSON(http.StatusOK, result)

}
