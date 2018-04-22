package main

import (
	"testing"
)

func TestDecode(t *testing.T) {
	data1 := &Decoder{
		Csr: "bad data",
	}

	test1, err1 := data1.Decode()
	if err1 == nil {
		t.Error("decode fail error")
	}
	if err1.Error() != "Csr decode error" {
		t.Error("decode error message is invalid")
	}
	if test1 != nil {
		t.Error("return value error on decode error statement")
	}

	data2 := &Decoder{
		Csr: `-----BEGIN CERTIFICATE REQUEST-----
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
-----END CERTIFICATE REQUEST-----`,
	}

	test2, err2 := data2.Decode()
	if err2 != nil {
		t.Error("decode error on valid string")
	}
	if test2 == nil {
		t.Error("parse faild at certificate request")
	}
	if test2.Subject.Country[0] != "JP" {
		t.Error("parse error on Subject")
	}
}
