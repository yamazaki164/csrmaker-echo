package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var e *echo.Echo

func newApp() {
	e = echo.New()
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

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
