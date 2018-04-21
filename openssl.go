package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
)

type OpenSsl struct {
	Csr *CsrParam
	Key *rsa.PrivateKey
	MarshalPKCS1PrivateKey []byte
}

func (x *OpenSsl) GeneratePrivateKey() ([]byte, error) {
	var err error
	x.Key, err = rsa.GenerateKey(rand.Reader, int(x.Csr.KeyBit))
	if err != nil {
		return nil, err
	}

	x.MarshalPKCS1PrivateKey = x509.MarshalPKCS1PrivateKey(x.Key)
	block := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x.MarshalPKCS1PrivateKey,
	}

	if x.Csr.EncryptCbc == Enctype_none {
		return pem.EncodeToMemory(block), nil
	}

	pass := []byte(x.Csr.PassPhrase)
	switch x.Csr.EncryptCbc {
		case Enctype_aes128:
			block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, pass, x509.PEMCipherAES128)
		case Enctype_aes192:
			block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, pass, x509.PEMCipherAES192)
		case Enctype_aes256:
			block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, pass, x509.PEMCipherAES256)
		case Enctype_des3:
			block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, pass, x509.PEMCipher3DES)
		default:
			err = errors.New("Encrypt CBC is not allowed")
	}

	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(block), nil
}

func (x *OpenSsl) GenerateCsr() ([]byte, error) {
	data := &x509.CertificateRequest{
		Subject: pkix.Name {
			Country: []string{x.Csr.Country},
			Province: []string{x.Csr.State},
			Locality: []string{x.Csr.Locality},
			Organization: []string{x.Csr.OrganizationalName},
			OrganizationalUnit: []string{x.Csr.OrganizationalUnit},
			CommonName: x.Csr.CommonName,
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
		Signature: x.MarshalPKCS1PrivateKey,
	}

	csr, err := x509.CreateCertificateRequest(rand.Reader, data, x.Key)
	if err != nil {
		return nil, err
	}

	block := &pem.Block {
		Type: "CERTIFICATE REQUEST",
		Bytes: csr,
	}

	return pem.EncodeToMemory(block), nil
}

func NewOpenSsl(param *CsrParam) *OpenSsl {
	s := &OpenSsl{
		Csr: param,
	}

	return s
}