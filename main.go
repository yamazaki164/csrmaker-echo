package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type TemplateRenderer struct {
	templates *template.Template
}

var e *echo.Echo

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e = echo.New()
	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}

	e.GET("/", indexHandler)
	e.POST("/create", createHandler)
	e.Logger.Fatal(e.Start(":9000"))
}
