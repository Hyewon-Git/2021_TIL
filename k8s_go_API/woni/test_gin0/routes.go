// routes.go

package main

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) *gin.Engine {
	/*
		//라우터 만들기 >> 다양한 url 처리를 위함
		router := gin.Default()
		//템플릿 로드
		router.LoadHTMLGlob("template/*")
	*/
	// Handle the index route
	router.GET("/", showIndexPage)

	//개별뉴스보는 링크
	router.GET("/article/view/:article_id", getArticle)

	return router
}
