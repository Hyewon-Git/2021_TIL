package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	//라우터 만들기 >> 다양한 url 처리를 위함
	router := gin.Default()
	//템플릿 로드
	router.LoadHTMLGlob("template/*")

	/*
		// 경로처리 정의 : 애플리케이션을 다양한 경로로 나누고 + 각 경로에 대한 핸들러 정의
		router.GET("/", func(c *gin.Context) { //c로 처리하는데 필요할수있는요청에대한 함모든정보 포
			c.HTML(
				http.StatusOK, //200ok로 HTTP상태설정
				"index.html",  //템플릿 연결
				gin.H{
					"title": "Homepage",
				},
			)
		})
	*/
	// 위에 코드대신 routes.go파일에 initializeRoutes함수에다 "/"주소에 showIndexPage로 링크걸음
	r := initializeRoutes(router)

	//시작
	r.Run()

}
