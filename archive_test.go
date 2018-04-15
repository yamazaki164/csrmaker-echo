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
		t.Errorf("nil pointer")
	}

	if !bytes.Equal(data["test"], test.targets["test"]) {
		t.Errorf("initialize error on targets")
	}

	if test.Buffer == nil {
		t.Errorf("buffer is nil")
	}
}

func TestCompress(t *testing.T) {
	data := map[string][]byte{}
	test := NewArchive(data)
	b := test.Buffer.Bytes()
	test.Compress()
	if test.Buffer == nil {
		t.Errorf("buffer is nil")
	}

	if bytes.Equal(test.Buffer.Bytes(), b) {
		t.Errorf("buffer is not compress")
	}
}
