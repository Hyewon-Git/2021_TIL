package main

import (
	"fmt"
	"k8s_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("here main")

	//라우터 만들기 >> 다양한 url 처리를 위함
	router := gin.Default()
	//템플릿 로드
	//router.LoadHTMLGlob("template/*")

	routes.InitializeRoutes(router)

	//시작
	router.Run()

}
