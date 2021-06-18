package main

import (
	"test0/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) *gin.Engine {
	/*
		//라우터 만들기 >> 다양한 url 처리를 위함
		router := gin.Default()
		//템플릿 로드
		router.LoadHTMLGlob("template/*")
	*/
	// 모든 item 조회
	router.GET("/", handlers.GetAllItem)
	// 특정 item pname으로 조회
	router.POST("/search", handlers.PostSearchItem)

	//item 추가
	router.POST("/post-item", handlers.PostItem)

	return router
}
