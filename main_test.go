package main

import (
	"testing"
)

func TestNewApp(t *testing.T) {
	if e != nil {
		t.Error("variable e is not nil on first")
	}
	newApp()
	if e == nil {
		t.Error("call newApp. but variable e was not initialzed.")
	}
}
