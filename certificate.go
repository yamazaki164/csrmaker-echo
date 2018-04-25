package main

import (
	"crypto/x509"
	"crypto/x509/pkix"
)

type Certificate struct {
	Issue     *Issue   `json:"issue"`
	Subject   *Subject `json:"subject"`
	NotAfter  string   `json:"notAfter"`
	NotBefore string   `json:"notBefore"`
}

func NewCertificate(cert *x509.Certificate) *Certificate {
	certificate := &Certificate{
		Issue:     NewIssueFromPkixName(&cert.Issuer),
		Subject:   NewSubjectFromPkixName(&cert.Subject),
		NotAfter:  cert.NotAfter.String(),
		NotBefore: cert.NotBefore.String(),
	}

	return certificate
}

type Issue struct {
	Country            string `json:"country"`
	OrganizationalName string `json:"organizationalName"`
	CommonName         string `json:"commonName"`
}

func NewIssueFromPkixName(x *pkix.Name) *Issue {
	issue := &Issue{
		Country:            x.Country[0],
		OrganizationalName: x.Organization[0],
		CommonName:         x.CommonName,
	}

	return issue
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
