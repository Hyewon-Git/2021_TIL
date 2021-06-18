package handlers

import (
	"fmt"
	"log"
	"net/http"
	"test0copy/models"

	"github.com/gin-gonic/gin"
)

// AddItem
func PostItem(c *gin.Context) {
	var newitem models.ItemStruct
	//var newitem models.ItemStruct
	err := c.Bind(&newitem) //  Post로입력받은 body값들을 newitem에 바인드시김
	if err != nil {
		log.Fatal(err)
		fmt.Println("bind error!")
	}
	fmt.Println(newitem)
	if newitem.PPRICE == 0 { //setting >> 가격은 0이없다!
		c.JSON(http.StatusOK, "Null Item: Add correct type Item")
	} else {
		fmt.Println(newitem) //제대롤 바인드됬는지 확인
		models.Setup()
		models.AddItem(newitem)

		models.Close()
		c.JSON(http.StatusOK, newitem)

	}

}

// AllItem
func GetAllItem(c *gin.Context) {
	var result gin.H
	fmt.Println("getallitem_start")
	models.Setup()
	itemlist := models.AllItem()

	if itemlist != nil {
		result = gin.H{
			"result": itemlist,
		}
	} else {
		result = gin.H{
			"result": nil,
		}
	}
	models.Close()
	c.JSON(http.StatusOK, result)
}

func PostSearchItem(c *gin.Context) {
	var searchPname models.ItemStruct
	err := c.Bind(&searchPname)
	if err != nil {
		log.Fatal(err)
		fmt.Println("bind error")
	}
	models.Setup()
	searchItems := models.SearchItem(searchPname.PNAME)
	models.Close()
	if searchItems == nil {
		c.JSON(http.StatusOK, "No searching item")
	} else {
		c.JSON(http.StatusOK, searchItems)
	}

}
