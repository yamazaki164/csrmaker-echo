package main

import (
	"github.com/labstack/echo"
)

var e *echo.Echo

func newApp() {
	e = echo.New()
	e.POST("/create", createHandler)
	e.POST("/csr/checker", doCsrCheckHandler)
	e.POST("/ssl/checker", doSslCheckHandler)

	e.Static("/", "public")
}

func main() {
	newApp()
	e.Logger.Fatal(e.Start(":9000"))
}
