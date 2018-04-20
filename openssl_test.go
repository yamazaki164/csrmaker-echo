package main

import (
	"os"
	"testing"
)

func beforeTest() {
	config = &Config{
		Cmd: `C:\OpenSSL-Win64\bin\openssl.exe`,
	}
}

func dummyCsr() CsrParam {
	p := CsrParam{
		EncryptCbc:         "des3",
		KeyBit:             2048,
		PassPhrase:         "test",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "meguro",
		OrganizationalName: "test company",
		OrganizationalUnit: "test unit",
		CommonName:         "test.example.com",
	}

	return p
}

func dummy() openssl {
	p := dummyCsr()
	o := openssl{
		csrParam: &p,
	}

	return o
}

func TestOpensslCmd(t *testing.T) {
	beforeTest()

	test := opensslCmd()
	if test != `C:\OpenSSL-Win64\bin\openssl.exe` {
		t.Errorf("opensslCmd error")
	}
}

func TestTmpPath(t *testing.T) {
	beforeTest()

	test := tmpPath()
	if test == "" {
		t.Errorf("path error")
	}
}

func TestCreateTempFile(t *testing.T) {
	beforeTest()

	test := CreateTempFile("test")
	_, e := os.Stat(test)
	if os.IsNotExist(e) {
		t.Errorf("temp file create error")
	}
}

func TestRemoveTempFile(t *testing.T) {
	beforeTest()
	test := CreateTempFile("test")
	RemoveTempFile(test)
	_, e := os.Stat(test)
	if os.IsExist(e) {
		t.Errorf("temp file remove error")
	}
}

func TestNewOpenssl(t *testing.T) {
	beforeTest()

	p := dummyCsr()
	o := NewOpenssl(&p)
	if o == nil {
		t.Errorf("nil pointer error")
	}
}

func TestRemoveKeyFile(t *testing.T) {
	beforeTest()

	test := openssl{
		keyFile: CreateTempFile("key"),
	}
	test.RemoveKeyFile()
	_, e := os.Stat(test.keyFile)
	if os.IsExist(e) {
		t.Errorf("key file remove error")
	}
}

func TestRemoveCsrFile(t *testing.T) {
	beforeTest()

	test := openssl{
		csrFile: CreateTempFile("csr"),
	}
	test.RemoveCsrFile()
	_, e := os.Stat(test.csrFile)
	if os.IsExist(e) {
		t.Errorf("csr file remove error")
	}
}

func TestGenrsaCommandOpt(t *testing.T) {
	beforeTest()
	o := dummy()

	test := o.GenrsaCommandOpt()
	if test[0] != "genrsa" {
		t.Errorf("[0] is not genrsa")
	}
	if test[1] != "-out" {
		t.Errorf("[1] is not -out")
	}
	if test[2] != "" {
		t.Errorf("[2] is not key")
	}
	if test[3] != "-des3" {
		t.Errorf("[3] is not encrypt cbc")
	}
	if test[4] != "-passout" {
		t.Errorf("[4] is not -passout")
	}
	if test[5] != "pass:test" {
		t.Errorf("[5] is not pass:test")
	}
	if test[6] != "2048" {
		t.Errorf("[6] is not 2048")
	}
}

func TestGenerateKey(t *testing.T) {
	beforeTest()

	o := dummy()
	o.keyFile = CreateTempFile("key")

	o.GenerateKey()

	f, e := os.Stat(o.keyFile)
	if os.IsNotExist(e) {
		t.Errorf("generate key error")
	}
	if f.Size() == 0 {
		t.Errorf("generate key error")
	}
}

func TestSetKeyRaw(t *testing.T) {
	beforeTest()

	o := dummy()
	o.keyFile = CreateTempFile("key")

	if len(o.KeyRaw) != 0 {
		t.Errorf("KeyRaw is not nil")
	}

	o.GenerateKey()
	o.SetKeyRaw()

	if len(o.KeyRaw) == 0 {
		t.Errorf("KeyRaw is empty")
	}

	defer func() {
		err := recover()
		if err != nil {
			t.Log("recover")
		}
	}()

	o = dummy()
	o.keyFile = "fuga"
	o.SetKeyRaw()
}

func TestSubj(t *testing.T) {
	beforeTest()

	o := dummy()

	test := o.subj()
	if test != "/C=JP/ST=Tokyo/L=meguro/O=test company/OU=test unit/CN=test.example.com" {
		t.Errorf("subj error")
	}
}

func TestReqNewCommandOpt(t *testing.T) {
	beforeTest()
	o := dummy()
	o.keyFile = CreateTempFile("key")
	o.csrFile = CreateTempFile("csr")

	test := o.ReqNewCommandOpt()
	if test[0] != "req" {
		t.Errorf("[0] is not req")
	}
	if test[1] != "-new" {
		t.Errorf("[1] is not -new")
	}
	if test[2] != "-sha256" {
		t.Errorf("[2] is not -sha256")
	}
	if test[3] != "-key" {
		t.Errorf("[3] is not -key")
	}
	if test[4] != o.keyFile {
		t.Errorf("[4] is not keyFile path")
	}
	if test[5] != "-subj" {
		t.Errorf("[5] is not -subj")
	}
	if test[6] != o.subj() {
		t.Errorf("[6] is not subj value")
	}
}

func TestGenerateCsr(t *testing.T) {
	beforeTest()
	o := dummy()
	o.keyFile = CreateTempFile("key")
	o.csrFile = CreateTempFile("csr")

	o.GenerateKey()
	o.GenerateCsr()

	f, e := os.Stat(o.csrFile)
	if os.IsNotExist(e) {
		t.Errorf("generate csr error")
	}
	if f.Size() == 0 {
		t.Errorf("generate csr error")
	}
}

func TestSetCsrRaw(t *testing.T) {
	beforeTest()
	o := dummy()
	o.keyFile = CreateTempFile("key")
	o.csrFile = CreateTempFile("csr")

	o.GenerateKey()
	if len(o.CsrRaw) != 0 {
		t.Errorf("csr is not nil")
	}
	o.GenerateCsr()
	o.SetCsrRaw()
	if len(o.CsrRaw) == 0 {
		t.Errorf("csr is empty")
	}

	defer func() {
		err := recover()
		if err != nil {
			t.Log("recover")
		}
	}()

	o = dummy()
	o.csrFile = "fuga"
	o.SetCsrRaw()
}
