package routes

import (
	"k8s_api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) *gin.Engine {

	router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	router.GET("/api/deployment/:namespace", handlers.GetDeployment)
	router.PUT("/api/deployment", handlers.PutDeployment)
	router.DELETE("/api/deployment", handlers.DeleteDeployment)

	router.GET("/api/service/:namespace", handlers.GetService)
	router.PUT("/api/service", handlers.PutService)
	router.DELETE("/api/service", handlers.DeleteService)

	return router
}
