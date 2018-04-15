package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yamazaki164/csrmaker-echo/model"
)

type TemplateRenderer struct {
	templates *template.Template
}

var config *Config
var e *echo.Echo

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func indexHandler(c echo.Context) error {
	keyBits := model.KeyBit
	encryptCbcs := model.EncryptCbc

	data := map[string]interface{}{
		"keyBits":     keyBits,
		"encryptCbcs": encryptCbcs,
	}
	return c.Render(http.StatusOK, "index.html", data)
}

func createHandler(c echo.Context) error {
	csr := &model.CsrParam{}
	if err := c.Bind(csr); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	//validation
	if isValid, errors := csr.Validate(); !isValid {
		return c.JSON(http.StatusBadRequest, errors)
	}

	s := NewOpenssl(csr)
	pass := ""
	if csr.EncryptCbc != model.Enctype_none {
		pass = csr.PassPhrase
	}
	files := map[string][]byte{
		"key.txt":  s.KeyRaw,
		"csr.txt":  s.CsrRaw,
		"pass.txt": []byte(pass),
	}
	ac := NewArchive(files)
	ac.Compress()

	return c.JSONBlob(http.StatusOK, ac.Buffer.Bytes())
}

func main() {
	e = echo.New()
	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	config = NewConfig()

	e.GET("/", indexHandler)
	e.POST("/create", createHandler)
	e.Logger.Fatal(e.Start(":9000"))
}
