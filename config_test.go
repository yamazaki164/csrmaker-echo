package main

import (
	"testing"
)

func TestConfigFilePath(t *testing.T) {
	test1 := ConfigFilePath()
	if test1 != "./config.toml" {
		t.Error("configFilePath is not ./config.toml")
	}
}

func TestNewConfig(t *testing.T) {
	configFile = "/test"
	test1 := NewConfig()
	if test1.Cmd != "/usr/bin/openssl" {
		t.Error("default Cmd is invalid")
	}

	configFile = "./test/config.toml"
	test2 := NewConfig()
	if test2.Cmd != "/path/to/command" {
		t.Error("decode error")
	}
}
