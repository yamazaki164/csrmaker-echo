package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func createHandler(c echo.Context) error {
	csr := &CsrParam{}
	if err := c.Bind(csr); err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorParam(err))
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
		return c.JSON(http.StatusInternalServerError, NewErrorParam(err))
	}
	CsrRaw, err := s.GenerateCsr()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorParam(err))
	}

	files := map[string][]byte{
		"key.txt":  KeyRaw,
		"csr.txt":  CsrRaw,
		"pass.txt": []byte(pass),
	}
	ac := NewArchive(files)
	if err := ac.Compress(); err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorParam(err))
	}

	return c.JSONBlob(http.StatusOK, ac.Buffer.Bytes())
}

func doCsrCheckHandler(c echo.Context) error {
	data := &CsrDecoder{}
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorParam(err))
	}

	x, err := data.Decode()
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorParam(err))
	}

	csr := NewCsrParamFromPkixName(&x.Subject)
	return c.JSON(http.StatusOK, csr)
}

func doSslCheckHandler(c echo.Context) error {
	data := &SslDecoder{}
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorParam(err))
	}

	x, err := data.Decode()
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorParam(err))
	}

	cert := NewCertificate(x)
	return c.JSON(http.StatusOK, cert)
}
