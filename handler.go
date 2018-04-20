package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func indexHandler(c echo.Context) error {
	keyBits := KeyBit
	encryptCbcs := EncryptCbc

	data := map[string]interface{}{
		"keyBits":     keyBits,
		"encryptCbcs": encryptCbcs,
	}
	return c.Render(http.StatusOK, "index.html", data)
}

func createHandler(c echo.Context) error {
	csr := &CsrParam{}
	if err := c.Bind(csr); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	//validation
	if isValid, errors := csr.Validate(); !isValid {
		return c.JSON(http.StatusBadRequest, errors)
	}

	s := NewOpenssl(csr)
	pass := ""
	if csr.EncryptCbc != Enctype_none {
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
