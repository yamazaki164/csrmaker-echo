package main

import (
	"errors"
	"testing"
)

func TestNewErrorParam(t *testing.T) {
	data1 := errors.New("data1")
	test1 := NewErrorParam(data1)

	if test1["error"] != "data1" {
		t.Error("ErrorParam: error field is invalid")
	}
}
