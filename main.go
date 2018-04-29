package main

import (
	"github.com/labstack/echo"
)

var e *echo.Echo

func newApp() {
	e = echo.New()
	g := e.Group("/api")
	g.POST("/create", createHandler)
	g.POST("/csr/checker", doCsrCheckHandler)
	g.POST("/ssl/checker", doSslCheckHandler)

	e.Static("/", "public")
}

func main() {
	newApp()
	e.Logger.Fatal(e.Start(":9000"))
}
