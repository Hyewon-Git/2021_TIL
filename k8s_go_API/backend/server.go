package main

import (
	"net/http"

	"example.com/soonbee/apiserver/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}   method=${method}, uri=${uri}, status=${status}\n",
	}))

	v1Group := e.Group("/api/v1")
	controller.JobController{}.Init(v1Group)

	e.Logger.Fatal(e.StartTLS(":1323", "private.crt", "private.key"))
}
