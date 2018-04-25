package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func TestIndexHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	indexHandler(c)
}

func TestCreateHandlerWithEmptyParam(t *testing.T) {
	e := echo.New()
	req1 := httptest.NewRequest(echo.POST, "/create", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/create")

	createHandler(c1)
	if rec1.Result().StatusCode != http.StatusBadRequest {
		t.Error("bad status code on empty params")
	}
}

func TestCreateHandlerWithInvalidParam(t *testing.T) {
	e := echo.New()
	param := make(url.Values)
	param.Set("keyBit", "2048")
	param.Set("encryptCbc", "aes128")
	param.Set("passPhrase", "")
	param.Set("country", "")
	param.Set("state", "")
	param.Set("locality", "")
	param.Set("organizationalName", "")
	param.Set("organizationalUnit", "")
	param.Set("commonName", "")

	req := httptest.NewRequest(echo.POST, "/create", strings.NewReader(param.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/create")
	err := createHandler(c)
	if err != nil {
		t.Error("createHandler return value error")
	}
	if rec.Result().StatusCode != http.StatusBadRequest {
		t.Error("bad status code on validate")
	}
}

func TestCreateHandlerWithValidParam(t *testing.T) {
	e := echo.New()
	param := make(url.Values)
	param.Set("keyBit", "2048")
	param.Set("encryptCbc", "aes128")
	param.Set("passPhrase", "1234")
	param.Set("country", "JP")
	param.Set("state", "Tokyo")
	param.Set("locality", "Meguro-Ku")
	param.Set("organizationalName", "test")
	param.Set("organizationalUnit", "test")
	param.Set("commonName", "test.example.com")

	req := httptest.NewRequest(echo.POST, "/create", strings.NewReader(param.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/create")
	err := createHandler(c)
	if err != nil {
		t.Error("createHandler return value error")
	}
	if rec.Result().StatusCode != http.StatusOK {
		t.Error("bad code error on validate with valid params")
	}
}

func TestCsrCheckerHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/csr/checker", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/csr-checker")

	csrCheckerHandler(c)
}

func TestDoCsrCheckHandlerWithEmptyParam(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/csr/checker", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/csr/checker")

	doCsrCheckHandler(c)
	if rec.Result().StatusCode != http.StatusBadRequest {
		t.Error("bad status code on empty params")
	}
}

func TestDoCsrCheckHandlerWithInvalidParam(t *testing.T) {
	e := echo.New()
	param := make(url.Values)
	param.Set("csr", "abc")

	req := httptest.NewRequest(echo.POST, "/csr/checker", strings.NewReader(param.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/csr/checker")

	doCsrCheckHandler(c)
	if rec.Result().StatusCode != http.StatusBadRequest {
		t.Error("bad status code on invalid params")
	}
}

func TestDoCsrCheckHandlerWithValidParam(t *testing.T) {
	e := echo.New()
	param := make(url.Values)
	param.Set("csr", `-----BEGIN CERTIFICATE REQUEST-----
MIICpTCCAY0CAQAwYDELMAkGA1UEBhMCSlAxDjAMBgNVBAgTBVRva3lvMRIwEAYD
VQQHEwlNZWd1cm8tS3UxDDAKBgNVBAoTA29ybzEJMAcGA1UECxMAMRQwEgYDVQQD
Ewt3d3cub3JvLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANuW
6Ve1Mx1EpvT09lH42sK9CJ1MkBk52b74dt/SKBt94kPN/wGfvJ+s52tDbmK6sysR
3LLTW0nFkKAT4epIJvPR3zKcjutsr4BqqaFbG3lW3Gdz/WykUGeQCm+1ucrx1lxm
cgFUqT/dXUTu5zgLlSyIek1GiKE756/RKZCMGSuI1S6qQoFJV3+nWKUJKjQf+rt9
0ihzDlYu3BcwobSj5LBNugZt0UniWFVMPYLyXY40j7dUBUu7G0f611WGEFxL2L1z
wdnjGd6W0NJbs0XQHsPI0kdLkA4JjFVolopNCyEsN0c3GYu10iAhpdfmQKoqUlLy
HpgJIAR2VbWpWa9WoFsCAwEAAaAAMA0GCSqGSIb3DQEBCwUAA4IBAQC6KYOUz36T
q26EAW+VKP7ezb2Y7+OMElHjIWJnA4RTfBr15tfgGg9yU3FztRKZVtZmCNgNcXek
iNvYho8T+0MrgZT1t613jNU9ZvnEpsM4i4qVJ+gUayULLppfp41s/fX29gWaMC6e
Nnbj7mf4U7WHl6bAOX2wtGAuUop5DpDq3rEJgUeRypAojh1a7XWITH8F/wga5w5/
Kx8CLMnwCXxqczZ+GPpmGJpv6puf90AZ0Bsef2JI52NuX4LObsRSYUPnZLgvzdby
STNvqhZTfhd9osOl4SXd+dF3/g2mrvTSM0y/iSEtfK8cQbXA3shOsPSu36ARJB9v
CDrKx9V2DqJC
-----END CERTIFICATE REQUEST-----`)

	req := httptest.NewRequest(echo.POST, "/csr/checker", strings.NewReader(param.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/csr/checker")

	doCsrCheckHandler(c)
	if rec.Result().StatusCode != http.StatusOK {
		t.Error("bad status code on valid params")
	}
}
