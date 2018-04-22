package main

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type Decoder struct {
	Csr string `json:"csr" form:"csr"`
}

func (d *Decoder) Decode() (*x509.CertificateRequest, error) {
	b := bytes.NewBufferString(d.Csr)
	p, _ := pem.Decode(b.Bytes())
	if p == nil {
		return nil, errors.New("Csr decode error")
	}

	return x509.ParseCertificateRequest(p.Bytes)
}
