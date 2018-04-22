package main

import (
	"regexp"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	ssl := &OpenSsl{
		Csr: &CsrParam{
			KeyBit:             2048,
			EncryptCbc:         "aes128",
			PassPhrase:         "test",
			Country:            "JP",
			State:              "Tokyo",
			Locality:           "test",
			OrganizationalName: "test",
			CommonName:         "example.com",
		},
	}

	test1, err1 := ssl.GeneratePrivateKey()
	if err1 != nil {
		t.Error(err1)
	}

	if match, e := regexp.MatchString("AES-128-CBC", string(test1)); e != nil {
		t.Error(e)
	} else if !match {
		t.Error("not match AES-128-CBC")
	}

	ssl.Csr.EncryptCbc = "aes192"
	test2, err2 := ssl.GeneratePrivateKey()
	if err2 != nil {
		t.Error(err2)
	}

	if match, e := regexp.MatchString("AES-192-CBC", string(test2)); e != nil {
		t.Error(e)
	} else if !match {
		t.Log(string(test2))
		t.Error("not match AES-192-CBC")
	}

	ssl.Csr.EncryptCbc = "aes256"
	test3, err3 := ssl.GeneratePrivateKey()
	if err3 != nil {
		t.Error(err3)
	}

	if match, e := regexp.MatchString("AES-256-CBC", string(test3)); e != nil {
		t.Error(e)
	} else if !match {
		t.Log(string(test3))
		t.Error("not match AES-256-CBC")
	}

	ssl.Csr.EncryptCbc = "des3"
	test4, err4 := ssl.GeneratePrivateKey()
	if err4 != nil {
		t.Error(err4)
	}

	if match, e := regexp.MatchString("DES-EDE3-CBC", string(test4)); e != nil {
		t.Error(e)
	} else if !match {
		t.Log(string(test4))
		t.Error("not match DES-EDE3-CBC")
	}

	ssl.Csr.EncryptCbc = "dummy"
	_, err5 := ssl.GeneratePrivateKey()
	if err5 != nil && err5.Error() != "Encrypt CBC is not allowed" {
		t.Error("error message is invalid: Encrypt CBC is not allowed")
	}

	ssl.Csr.KeyBit = 1
	_, err6 := ssl.GeneratePrivateKey()
	if err6 == nil {
		t.Error("KeyBit is invalid")
	}
}

func TestGenerateCsr(t *testing.T) {
	ssl := &OpenSsl{
		Csr: &CsrParam{
			KeyBit:             2048,
			EncryptCbc:         "aes128",
			PassPhrase:         "test",
			Country:            "JP",
			State:              "Tokyo",
			Locality:           "test",
			OrganizationalName: "test",
			CommonName:         "example.com",
		},
	}

	ssl.GeneratePrivateKey()
	_, e := ssl.GenerateCsr()
	if e != nil {
		t.Error(e)
	}
}

func TestNewOpenSsl(t *testing.T) {
	csr := &CsrParam{
		KeyBit:             2048,
		EncryptCbc:         "aes128",
		PassPhrase:         "test",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}

	test := NewOpenSsl(csr)
	if test.Csr != csr {
		t.Error("csr param error")
	}
}
