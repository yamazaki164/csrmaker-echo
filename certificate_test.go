package main

import (
	"crypto/x509/pkix"
	"testing"
)

func TestNewCertificate(t *testing.T) {
	test1 := NewCertificate()
}

func TestNewIssueFromPkixName(t *testing.T) {
	data1 := &pkix.Name{
		Country:      []string{"JP"},
		Organization: []string{"organization"},
		CommonName:   "test",
	}
	test1 := NewIssueFromPkixName(data1)
	if test1.Country != "JP" {
		t.Error("Country is invalid")
	}
	if test1.Organization != "organization" {
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
	if test1.OrganizationName != "O" {
		t.Error("OrganizationName is invalid")
	}
	if test1.OrganizationalUnit != "OU" {
		t.Error("OrganizationalUnit is invalid")
	}
	if test1.CommonName != "test" {
		t.Error("CommonName is invalid")
	}
}
