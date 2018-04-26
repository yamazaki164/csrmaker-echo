package main

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"testing"
	"time"
)

func TestNewCertificate(t *testing.T) {
	data1 := &x509.Certificate{
		Issuer: pkix.Name{
			Country:      []string{"JP"},
			Organization: []string{"organization"},
			CommonName:   "test",
		},
		Subject: pkix.Name{
			Country:            []string{"JP"},
			Province:           []string{"S"},
			Locality:           []string{"locality"},
			Organization:       []string{"O"},
			OrganizationalUnit: []string{"OU"},
			CommonName:         "test",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now(),
	}
	test1 := NewCertificate(data1)
	if test1.Issuer == nil {
		t.Error("Issue is nil")
	}
	if test1.Subject == nil {
		t.Error("Subject is nil")
	}
	if test1.NotBefore == "" {
		t.Error("NotBefore is empty")
	}
	if test1.NotAfter == "" {
		t.Error("NotAfter is empty")
	}
}

func TestNewIssuerFromPkixName(t *testing.T) {
	data1 := &pkix.Name{
		Country:      []string{"JP"},
		Organization: []string{"organization"},
		CommonName:   "test",
	}
	test1 := NewIssuerFromPkixName(data1)
	if test1.Country != "JP" {
		t.Error("Country is invalid")
	}
	if test1.OrganizationalName != "organization" {
		t.Error("Organization is invalid")
	}
	if test1.CommonName != "test" {
		t.Error("CommonName is invalid")
	}
}

func TestNewSubjectFromPkixName(t *testing.T) {
	data1 := &pkix.Name{
		Country:            []string{"JP"},
		Province:           []string{"S"},
		Locality:           []string{"locality"},
		Organization:       []string{"O"},
		OrganizationalUnit: []string{"OU"},
		CommonName:         "test",
	}

	test1 := NewSubjectFromPkixName(data1)
	if test1.Country != "JP" {
		t.Error("Country is invalid")
	}
	if test1.State != "S" {
		t.Error("State is invalid")
	}
	if test1.Locality != "locality" {
		t.Error("Locality is invalid")
	}
	if test1.OrganizationalName != "O" {
		t.Error("OrganizationName is invalid")
	}
	if test1.OrganizationalUnit != "OU" {
		t.Error("OrganizationalUnit is invalid")
	}
	if test1.CommonName != "test" {
		t.Error("CommonName is invalid")
	}
}
