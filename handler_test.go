package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func TestCreateHandlerWithEmptyParam(t *testing.T) {
	e := echo.New()
	req1 := httptest.NewRequest(echo.POST, "/api/create", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/api/create")

	createHandler(c1)
	if rec1.Result().StatusCode != http.StatusBadRequest {
		t.Error("createHandler: bad status code on empty params")
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

	req := httptest.NewRequest(echo.POST, "/api/create", strings.NewReader(param.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/create")
	err := createHandler(c)
	if err != nil {
		t.Error("createHandler: createHandler return value error")
	}
	if rec.Result().StatusCode != http.StatusBadRequest {
		t.Error("createHandler: bad status code on validate")
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

	req := httptest.NewRequest(echo.POST, "/api/create", strings.NewReader(param.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/create")
	err := createHandler(c)
	if err != nil {
		t.Error("createHandler: createHandler return value error")
	}
	if rec.Result().StatusCode != http.StatusOK {
		t.Error("createHandler: bad code error on validate with valid params")
	}
}

func TestDoCsrCheckHandlerWithEmptyParam(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/api/csr/checker", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/csr/checker")

	doCsrCheckHandler(c)
	if rec.Result().StatusCode != http.StatusBadRequest {
		t.Error("doCsrCheckHandler: bad status code on empty params")
	}
}

func TestDoCsrCheckHandlerWithInvalidParam(t *testing.T) {
	e := echo.New()
	param := make(url.Values)
	param.Set("csr", "abc")

	req := httptest.NewRequest(echo.POST, "/api/csr/checker", strings.NewReader(param.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/csr/checker")

	doCsrCheckHandler(c)
	if rec.Result().StatusCode != http.StatusBadRequest {
		t.Error("doCsrCheckHandler: bad status code on invalid params")
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

	req := httptest.NewRequest(echo.POST, "/api/csr/checker", strings.NewReader(param.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/csr/checker")

	doCsrCheckHandler(c)
	if rec.Result().StatusCode != http.StatusOK {
		t.Error("doCsrCheckHandler: bad status code on valid params")
	}
}

func TestDoSslCheckHandlerWithEmptyParam(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/api/ssl/checker", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ssl/checker")

	doSslCheckHandler(c)
	if rec.Result().StatusCode != http.StatusBadRequest {
		t.Error("doSslCheckHandler: bad status code on empty params")
	}
}

func TestDoSslCheckHandlerWithInvalidParam(t *testing.T) {
	e := echo.New()
	param := make(url.Values)
	param.Set("certificate", "abc")

	req := httptest.NewRequest(echo.POST, "/api/ssl/checker", strings.NewReader(param.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ssl/checker")

	doSslCheckHandler(c)
	if rec.Result().StatusCode != http.StatusBadRequest {
		t.Error("doSslCheckHandler: bad status code on invalid params")
	}
}

func TestDoSslCheckHandlerWithValidParam(t *testing.T) {
	e := echo.New()
	param := make(url.Values)
	param.Set("certificate", `-----BEGIN CERTIFICATE-----
MIIDajCCAlICCQCAeI0i7y+VcjANBgkqhkiG9w0BAQsFADB3MQswCQYDVQQGEwJK
UDERMA8GA1UECAwISG9ra2FpZG8xEDAOBgNVBAcMB1NhcHBvcm8xFTATBgNVBAoM
DEV4YW1wbGUgSU5DLjEWMBQGA1UECwwNSVQgRGVwYXJ0bWVudDEUMBIGA1UEAwwL
ZXhhbXBsZS5jb20wHhcNMTgwNDI2MTQyMjI5WhcNMjgwNDIzMTQyMjI5WjB3MQsw
CQYDVQQGEwJKUDERMA8GA1UECAwISG9ra2FpZG8xEDAOBgNVBAcMB1NhcHBvcm8x
FTATBgNVBAoMDEV4YW1wbGUgSU5DLjEWMBQGA1UECwwNSVQgRGVwYXJ0bWVudDEU
MBIGA1UEAwwLZXhhbXBsZS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
AoIBAQDIQrzTFOPTr6MTAqhakY+0RT8OLi+09Sdlv42z2zwZ9/HFLukLh7PaHHNp
hEjbxrrqGO++sPEMjKE2VSS2391UUBpmQ5vHz7Y69m2rzWpP01m1Fl/oruJrOdY0
awz3+XHIrLCOscBmJ4+DZR1xJHwPG5rd706vm7sfvIJ/7/S8A8/0d6o0Xna9Fo8t
ZT1//8/8I9Ry8fCvNjdAmDoJdkvRN5WMFaChBZ2stWr9I4MesuKrGvyWCO4J1Lq0
9YHDorSzZNl7u4sw+e1iTv5NLIso4zza73VjxV+p5RJ2JsQEJPFeeE/Ldsox88Lo
CC5umStPNCtibkSfc2TAVV025oaDAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAJty
SnMN8+Ijup4IL6qkL4AhLlDRtU/yl2s3oXN8p0r1dZlbBbCjnyjsqImk/NnjxiYa
L3cQT5OXOhNA/o1iExkeMmvLlRhHLIYkqg6NVs2/KlrckBUdrrRnRkNHKzMy+Pi2
u7RfBCeVIF2Bm+MTZmRX2Rfh2M+hV6qdgrQVHF/hBFA17LAS5W3XBnLTTCgxr2nL
mIxGUOrZ1kZ0FTqtrTKTLOIKuWhN6FnAEC7TvbvZGK3D6B8YtlFQIarTTY3o4R3F
jN6pHmvDGiyDfJmCwykgiee+PPPYZYCo2xqmFd7wKGhtKuHw93ToaLTvxe3yDJla
aVeHNwXryBZyN3AXOP8=
-----END CERTIFICATE-----`)

	req := httptest.NewRequest(echo.POST, "/api/ssl/checker", strings.NewReader(param.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ssl/checker")

	doSslCheckHandler(c)
	if rec.Result().StatusCode != http.StatusOK {
		t.Error("doSslCheckHandler: bad status code on valid params")
	}
}
