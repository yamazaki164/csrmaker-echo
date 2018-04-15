package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Cmd string
}

func NewConfig() *Config {
	c := &Config{}
	if _, err := toml.DecodeFile("./config.toml", c); err != nil {
		return &Config{
			Cmd: "/usr/bin/openssl",
		}
	}
	
	return c
}
