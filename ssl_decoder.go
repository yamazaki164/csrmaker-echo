package main

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type SslDecoder struct {
	Certificate string `json:"certificate" form:"certificate"`
}

func (s *SslDecoder) Decode() (*x509.Certificate, error) {
	b := bytes.NewBufferString(s.Certificate)
	p, _ := pem.Decode(b.Bytes())
	if p == nil {
		return nil, errors.New("Certificate decode error")
	}

	return x509.ParseCertificate(p.Bytes)
}
