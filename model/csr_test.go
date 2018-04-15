package model

import (
	"testing"
)

func TestValidate(t *testing.T) {
	p1 := &CsrParam{
		KeyBit:             2048,
		EncryptCbc:         "aes128",
		PassPhrase:         "test",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}
	
	if v1, e1 := p1.Validate(); !v1 {
		t.Error(e1)
	}

	p2 := &CsrParam{
		KeyBit:             2048,
		EncryptCbc:         "aes128",
		PassPhrase:         "",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}
	
	if v2, e2 := p2.Validate(); v2 {
		t.Errorf("has error on empty PassPhrase")
		t.Error(e2)
	}

	p3 := &CsrParam{
		KeyBit:             2048,
		EncryptCbc:         "none",
		PassPhrase:         "",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}
	
	if v3, e3 := p3.Validate(); !v3 {
		t.Error(e3)
	}
}

func TestKeyBitOfValidate(t *testing.T) {
	message := "KeyBit is either 2048 or 4096"
	p1 := &CsrParam{
		KeyBit:             1028,
		EncryptCbc:         "aes128",
		PassPhrase:         "test",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}
	
	v1, e1 := p1.Validate()
	if v1 {
		t.Errorf("has no error")
	}
	
	if e1.KeyBit != message {
		t.Errorf("KeyBit validation error")
	}

	p2 := &CsrParam{
		KeyBit:             2048,
		EncryptCbc:         "aes128",
		PassPhrase:         "test",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}
	v2, e2 := p2.Validate()
	
	if !v2 {
		t.Error(e2)
	}

	p3 := &CsrParam{
		KeyBit:             4096,
		EncryptCbc:         "aes128",
		PassPhrase:         "test",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}
	v3, e3 := p3.Validate()
	if !v3 {
		t.Error(e3)
	}

	p4 := &CsrParam{
		KeyBit:             4097,
		EncryptCbc:         "aes128",
		PassPhrase:         "test",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}
	v4, e4 := p4.Validate()
	if v4 {
		t.Errorf("has no error on KeyBit")
	}
	if e4.KeyBit != message {
		t.Errorf("KeyBit validation error")
	}
}
