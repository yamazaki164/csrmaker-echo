package main

import (
	"crypto/x509"
	"crypto/x509/pkix"
)

type Certificate struct {
	Issuer    *Issuer  `json:"issuer"`
	Subject   *Subject `json:"subject"`
	NotAfter  string   `json:"notAfter"`
	NotBefore string   `json:"notBefore"`
}

func NewCertificate(cert *x509.Certificate) *Certificate {
	certificate := &Certificate{
		Issuer:    NewIssuerFromPkixName(&cert.Issuer),
		Subject:   NewSubjectFromPkixName(&cert.Subject),
		NotAfter:  cert.NotAfter.String(),
		NotBefore: cert.NotBefore.String(),
	}

	return certificate
}

type Issuer struct {
	Country            string `json:"country"`
	OrganizationalName string `json:"organizationalName"`
	CommonName         string `json:"commonName"`
}

func NewIssuerFromPkixName(x *pkix.Name) *Issuer {
	issuer := &Issuer{
		Country:            x.Country[0],
		OrganizationalName: x.Organization[0],
		CommonName:         x.CommonName,
	}

	return issuer
}

type Subject struct {
	Country            string `json:"country"`
	State              string `json:"state"`
	Locality           string `json:"locality"`
	OrganizationalName string `json:"organizationalName"`
	OrganizationalUnit string `json:"organizationalUnit"`
	CommonName         string `json:"commonName"`
}

func NewSubjectFromPkixName(x *pkix.Name) *Subject {
	subject := &Subject{
		Country:            x.Country[0],
		State:              x.Province[0],
		Locality:           x.Locality[0],
		OrganizationalName: x.Organization[0],
		OrganizationalUnit: x.OrganizationalUnit[0],
		CommonName:         x.CommonName,
	}

	return subject
}
