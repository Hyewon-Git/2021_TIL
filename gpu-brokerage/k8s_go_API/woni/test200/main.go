package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("<DOMAIN>")
	// Cache certificates
	// e.AutoTLSManager.Cache = autocert.DirCache("/usr/local/opt/openssl/bin/.keystone")
	// e.Use(middleware.Recover())
	// e.Use(middleware.Logger())

	// e.Pre(middleware.HTTPSRedirect())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to HYEWON-PAGE!</h1>
			<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
		`)
	})
	e.Logger.Fatal(e.StartTLS(":443", "private.crt", "private.key"))
}
