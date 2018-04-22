package main

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

func TestEncryptCbc(t *testing.T) {
	p1 := &CsrParam{
		KeyBit:             4096,
		EncryptCbc:         "aes",
		PassPhrase:         "test",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}

	v1, e1 := p1.Validate()
	if v1 {
		t.Error("validation error on EncryptCbc")
	}
	if e1.EncryptCbc != "EncryptCbc is aes128, aes192, aes256, des3, or none" {
		t.Error("validation message error on EncryptCbc")
	}
}

func TestPassPhrase(t *testing.T) {
	p1 := &CsrParam{
		KeyBit:             4096,
		EncryptCbc:         "des3",
		PassPhrase:         "",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}

	v1, e1 := p1.Validate()
	if v1 {
		t.Error("validation error on PassPhrase")
	}
	if e1.PassPhrase != "PassPhrase required" {
		t.Error("valation message error on PassPhrase")
	}

	p2 := &CsrParam{
		KeyBit:             4096,
		EncryptCbc:         "aes128",
		PassPhrase:         "",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}

	v2, e2 := p2.Validate()
	if v2 {
		t.Error("validation error on PassPhrase")
	}
	if e2.PassPhrase != "PassPhrase required" {
		t.Error("valation message error on PassPhrase")
	}

	p3 := &CsrParam{
		KeyBit:             4096,
		EncryptCbc:         "aes192",
		PassPhrase:         "",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}

	v3, e3 := p3.Validate()
	if v3 {
		t.Error("validation error on PassPhrase")
	}
	if e3.PassPhrase != "PassPhrase required" {
		t.Error("valation message error on PassPhrase")
	}

	p4 := &CsrParam{
		KeyBit:             4096,
		EncryptCbc:         "aes256",
		PassPhrase:         "",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}

	v4, e4 := p4.Validate()
	if v4 {
		t.Error("validation error on PassPhrase")
	}
	if e4.PassPhrase != "PassPhrase required" {
		t.Error("valation message error on PassPhrase")
	}
}

func TestCountry(t *testing.T) {
	p1 := &CsrParam{
		KeyBit:             4096,
		EncryptCbc:         "none",
		PassPhrase:         "",
		Country:            "JPN",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}

	v1, e1 := p1.Validate()
	if v1 {
		t.Error("validation error on Country")
	}
	if e1.Country != "Country must be 2 letter" {
		t.Error("valation message error on Country")
	}

	p2 := &CsrParam{
		KeyBit:             4096,
		EncryptCbc:         "aes128",
		PassPhrase:         "",
		Country:            "jp",
		State:              "Tokyo",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}

	v2, e2 := p2.Validate()
	if v2 {
		t.Error("validation error on Country")
	}
	if e2.Country != "Country must be 2 letter" {
		t.Error("valation message error on Country")
	}
}

func TestState(t *testing.T) {
	p1 := &CsrParam{
		KeyBit:             4096,
		EncryptCbc:         "none",
		PassPhrase:         "",
		Country:            "JP",
		State:              "",
		Locality:           "test",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}

	v1, e1 := p1.Validate()
	if v1 {
		t.Error("validation error on State")
	}
	if e1.State != "State required" {
		t.Error("valation message error on State")
	}
}

func TestLocality(t *testing.T) {
	p1 := &CsrParam{
		KeyBit:             4096,
		EncryptCbc:         "none",
		PassPhrase:         "",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "",
		OrganizationalName: "test",
		CommonName:         "example.com",
	}

	v1, e1 := p1.Validate()
	if v1 {
		t.Error("validation error on Locality")
	}
	if e1.Locality != "Locality required" {
		t.Error("valation message error on Locality")
	}
}

func TestOrganizationalName(t *testing.T) {
	p1 := &CsrParam{
		KeyBit:             4096,
		EncryptCbc:         "none",
		PassPhrase:         "",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "Meguro-ku",
		OrganizationalName: "",
		CommonName:         "example.com",
	}

	v1, e1 := p1.Validate()
	if v1 {
		t.Error("validation error on OrganizationalName")
	}
	if e1.OrganizationalName != "OrganizationalName required" {
		t.Error("valation message error on OrganizationalName")
	}
}

func TestCommonName(t *testing.T) {
	p1 := &CsrParam{
		KeyBit:             4096,
		EncryptCbc:         "none",
		PassPhrase:         "",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "Meguro-ku",
		OrganizationalName: "test",
		CommonName:         "",
	}

	v1, e1 := p1.Validate()
	if v1 {
		t.Error("validation error on CommonName")
	}
	if e1.CommonName != "CommonName required" {
		t.Error("valation message error on CommonName")
	}
}
