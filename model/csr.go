package model

import (
	"regexp"
)

const (
	Enctype_none   = "none"
	Enctype_des3   = "des3"
	Enctype_aes256 = "aes256"
	Enctype_aes192 = "aes192"
	Enctype_aes128 = "aes128"
)

var (
	EncryptCbc map[string]string = map[string]string{
		Enctype_none:   Enctype_none,
		Enctype_des3:   Enctype_des3,
		Enctype_aes256: Enctype_aes256,
		Enctype_aes192: Enctype_aes192,
		Enctype_aes128: Enctype_aes128,
	}

	KeyBit map[string]string = map[string]string{
		"2048": "2048",
		"4096": "4096",
	}
)

type CsrParam struct {
	KeyBit             uint   `json:"keyBit" form:"keyBit"`
	EncryptCbc         string `json:"encryptCbc" form:"encryptCbc"`
	PassPhrase         string `json:"passPhrase" form:"passPhrase"`
	Country            string `json:"country" form:"country"`
	State              string `json:"state" form:"state"`
	Locality           string `json:"locality" form:"locality"`
	OrganizationalName string `json:"organizationalName" form:"organizationalName"`
	OrganizationalUnit string `json:"organizationalUnit" form:"organizationalUnit"`
	CommonName         string `json:"commonName" form:"commonName"`
}

type CsrErrors struct {
	KeyBit             string `json:"keyBit" form:"keyBit"`
	EncryptCbc         string `json:"encryptCbc" form:"encryptCbc"`
	PassPhrase         string `json:"passPhrase" form:"passPhrase"`
	Country            string `json:"country" form:"country"`
	State              string `json:"state" form:"state"`
	Locality           string `json:"locality" form:"locality"`
	OrganizationalName string `json:"organizationalName" form:"organizationalName"`
	OrganizationalUnit string `json:"organizationalUnit" form:"organizationalUnit"`
	CommonName         string `json:"commonName" form:"commonName"`
}

func (c *CsrParam) Validate() (bool, *CsrErrors) {
	errors := &CsrErrors{}
	isValid := true

	if c.KeyBit != 2048 && c.KeyBit != 4096 {
		errors.KeyBit = "KeyBit is either 2048 or 4096"
		isValid = false
	}

	if !regexp.MustCompile("^(aes128|aes192|aes256|des3|none)$").MatchString(c.EncryptCbc) {
		errors.EncryptCbc = "EncryptCbc is aes128, aes192, aes256, des3, or none"
		isValid = false
	}

	if c.EncryptCbc != Enctype_none && c.PassPhrase == "" {
		errors.PassPhrase = "PassPhrase required"
		isValid = false
	}

	if !regexp.MustCompile("^[A-Z]{2}$").MatchString(c.Country) {
		errors.Country = "Country must be 2 letter"
		isValid = false
	}

	if c.State == "" {
		errors.State = "State required"
		isValid = false
	}
	if c.Locality == "" {
		errors.Locality = "Locality required"
		isValid = false
	}

	if c.OrganizationalName == "" {
		errors.OrganizationalName = "OrganizationalName required"
		isValid = false
	}

	if c.CommonName == "" {
		errors.CommonName = "CommonName required"
		isValid = false
	}

	return isValid, errors
}
