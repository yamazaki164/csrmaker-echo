package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Cmd string
}

var (
	configFile     string = "./config.toml"
	defaultCommand string = "/usr/bin/openssl"
)

func ConfigFilePath() string {
	return configFile
}

func NewConfig() *Config {
	c := &Config{}
	if _, err := toml.DecodeFile(ConfigFilePath(), c); err != nil {
		return &Config{
			Cmd: defaultCommand,
		}
	}

	return c
}
