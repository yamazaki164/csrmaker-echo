package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/yamazaki164/csrmaker-echo/model"
)

type openssl struct {
	csrParam *model.CsrParam
	keyFile  string
	KeyRaw   []byte
	csrFile  string
	CsrRaw   []byte
}

func opensslCmd() string {
	return config.Cmd
}

func tmpPath() string {
	return os.TempDir()
}

func CreateTempFile(prefix string) string {
	tmpfile, err := ioutil.TempFile(tmpPath(), prefix)
	if err != nil {
		panic(err)
	}
	defer tmpfile.Close()
	return tmpfile.Name()
}

func RemoveTempFile(file string) {
	if err := os.Remove(file); err != nil {
		panic(err)
	}
}

func NewOpenssl(csrParam *model.CsrParam) *openssl {
	o := &openssl{
		csrParam: csrParam,
		keyFile:  CreateTempFile("key"),
		csrFile:  CreateTempFile("csr"),
	}

	o.GenerateKey()
	o.GenerateCsr()

	defer o.RemoveCsrFile()
	defer o.RemoveKeyFile()

	o.SetKeyRaw()
	o.SetCsrRaw()

	return o
}

func (o *openssl) RemoveKeyFile() {
	RemoveTempFile(o.keyFile)
}

func (o *openssl) RemoveCsrFile() {
	RemoveTempFile(o.csrFile)
}

func (o *openssl) GenrsaCommandOpt() []string {
	opt := []string{
		"genrsa",
		"-out",
		o.keyFile,
	}

	if o.csrParam.EncryptCbc != model.Enctype_none {
		opt = append(opt, []string{
			"-" + o.csrParam.EncryptCbc,
			"-passout",
			"pass:" + o.csrParam.PassPhrase,
		}...)
	}

	opt = append(opt, strconv.FormatUint(uint64(o.csrParam.KeyBit), 10))

	return opt
}

func (o *openssl) GenerateKey() {
	ret, err := exec.Command(opensslCmd(), o.GenrsaCommandOpt()...).CombinedOutput()
	log.Println(ret)
	if err != nil {
		panic(err)
	}
}

func (o *openssl) SetKeyRaw() {
	var err error
	o.KeyRaw, err = ioutil.ReadFile(o.keyFile)
	if err != nil {
		panic(err)
	}
}

func (o *openssl) subj() string {
	return "/C=" + o.csrParam.Country + "/ST=" + o.csrParam.State + "/L=" + o.csrParam.Locality + "/O=" + o.csrParam.OrganizationalName + "/OU=" + o.csrParam.OrganizationalUnit + "/CN=" + o.csrParam.CommonName
}

func (o *openssl) ReqNewCommandOpt() []string {
	opt := []string{
		"req",
		"-new",
		"-sha256",
		"-key",
		o.keyFile,
		"-subj",
		o.subj(),
	}

	if o.csrParam.EncryptCbc != model.Enctype_none {
		opt = append(opt, []string{"-passin", "pass:" + o.csrParam.PassPhrase}...)
	}

	opt = append(opt, []string{"-out", o.csrFile}...)

	return opt
}

func (o *openssl) GenerateCsr() {
	ret, err := exec.Command(opensslCmd(), o.ReqNewCommandOpt()...).CombinedOutput()
	log.Println(ret)
	if err != nil {
		panic(err)
	}
}

func (o *openssl) SetCsrRaw() {
	var err error
	o.CsrRaw, err = ioutil.ReadFile(o.csrFile)
	if err != nil {
		panic(err)
	}
}
