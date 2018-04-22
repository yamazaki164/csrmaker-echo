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
	if isValid, err := csr.Validate(); !isValid {
		return c.JSON(http.StatusBadRequest, err)
	}

	s := NewOpenSsl(csr)
	pass := ""
	if csr.EncryptCbc != Enctype_none {
		pass = csr.PassPhrase
	}
	KeyRaw, err := s.GeneratePrivateKey()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	CsrRaw, err := s.GenerateCsr()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	files := map[string][]byte{
		"key.txt":  KeyRaw,
		"csr.txt":  CsrRaw,
		"pass.txt": []byte(pass),
	}
	ac := NewArchive(files)
	ac.Compress()

	return c.JSONBlob(http.StatusOK, ac.Buffer.Bytes())
}

func checkerHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "checker.html", nil)
}

func doCheckHandler(c echo.Context) error {
	data := &Decoder{}
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	x, err := data.Decode()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	csr := &CsrParam{
		Country:            x.Subject.Country[0],
		State:              x.Subject.Province[0],
		Locality:           x.Subject.Locality[0],
		OrganizationalName: x.Subject.Organization[0],
		OrganizationalUnit: x.Subject.OrganizationalUnit[0],
		CommonName:         x.Subject.CommonName,
	}

	return c.JSON(http.StatusOK, csr)
}
