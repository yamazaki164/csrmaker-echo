package main

import (
	"bytes"
	"testing"
)

func TestNewArchive(t *testing.T) {
	data := map[string][]byte{
		"test": []byte("test"),
	}

	test := NewArchive(data)
	if test == nil {
		t.Error("nil pointer")
	}

	if !bytes.Equal(data["test"], test.targets["test"]) {
		t.Error("initialize error on targets")
	}

	if test.Buffer == nil {
		t.Error("buffer is nil")
	}
}

func TestCompress(t *testing.T) {
	data := map[string][]byte{}
	test := NewArchive(data)
	b := test.Buffer.Bytes()
	err := test.Compress()
	if err != nil {
		t.Errorf("Compress: return err with %s", err.Error())
	}

	if test.Buffer == nil {
		t.Error("buffer is nil")
	}

	if bytes.Equal(test.Buffer.Bytes(), b) {
		t.Error("buffer is not compress")
	}
}

func TestCompressWithError(t *testing.T) {
	data := map[string][]byte{
		"fuga": nil,
	}
	test := &Archive{
		targets: data,
		Buffer:  nil,
	}

	err := test.Compress()
	if err == nil {
		t.Error("Compress: error on nil buffer")
	}
	if err.Error() != "Buffer nil pointer error" {
		t.Error("Compress: invalid error message at nil buffer")
	}
}
