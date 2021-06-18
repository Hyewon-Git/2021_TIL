package main

import "github.com/gin-gonic/gin"

func main() {
	//라우터 만들기 >> 다양한 url 처리를 위함
	router := gin.Default()
	//템플릿 로드
	//router.LoadHTMLGlob("template/*")

	r := InitializeRoutes(router)

	//시작
	r.Run()
}
