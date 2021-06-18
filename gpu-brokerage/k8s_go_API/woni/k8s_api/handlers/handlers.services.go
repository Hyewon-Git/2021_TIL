package handlers

import (
	"fmt"
	"k8s_api/api"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetService(c *gin.Context) {
	np := c.Param("namespace")
	fmt.Println(np)
	if np == "all" {
		np = ""
		//c.JSON(http.StatusBadRequest, "Please request namespace")
		//np = apiv1.NamespaceDefault
	}
	config := api.ConfirmConfig()
	clientset := api.SetClient(config)
	result := api.ListService(clientset, np)

	c.JSON(http.StatusOK, result)
}

func PutService(c *gin.Context) {
	var createSv api.ServiceStruct
	//var newitem models.ItemStruct
	err := c.Bind(&createSv) //  Put으로입력받은 body값들을 newitem에 바인드시김
	if err != nil {
		log.Fatal(err)
		fmt.Println("bind error!")
	}

	config := api.ConfirmConfig()
	clientset := api.SetClient(config)
	api.CreateService(clientset, createSv)

	//fmt.Printf("Created deployment %q.\n", dpName)
	c.JSON(http.StatusOK, createSv)

}

func DeleteService(c *gin.Context) {
	var deleteSv api.ServiceStruct
	//var newitem models.ItemStruct
	err := c.Bind(&deleteSv) //  Put으로입력받은 body값들을 newitem에 바인드시김
	if err != nil {
		log.Fatal(err)
		fmt.Println("bind error!")
	}
	config := api.ConfirmConfig()
	clientset := api.SetClient(config)
	api.DeleteService(clientset, deleteSv)

	result := fmt.Sprintf("Delete namespace : %s Service : %s", deleteSv.SNS, deleteSv.SNAME)
	c.JSON(http.StatusOK, result)

}
