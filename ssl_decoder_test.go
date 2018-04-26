package main

import (
	"testing"
)

func TestSslDecode(t *testing.T) {
	data1 := &SslDecoder{
		Certificate: "bad data",
	}

	test1, err1 := data1.Decode()
	if err1 == nil {
		t.Error("decode fail error")
	}
	if err1.Error() != "Certificate decode error" {
		t.Error("decode error message is invalid")
	}
	if test1 != nil {
		t.Error("return value error on decode error statement")
	}

	data2 := &SslDecoder{
		Certificate: `-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----`,
	}

	test2, err2 := data2.Decode()
	if err2 != nil {
		t.Error("decode error on valid string")
	}
	if test2 == nil {
		t.Error("parse faild at certificate")
	}
	if test2.Issuer.Organization[0] != "Example INC." || test2.Issuer.Country[0] != "JP" {
		t.Error("Issuer is invalid in test data")
	}
}
