package main

import (
	"AvitoTask/configs"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

const (
	configsPath = "confgigs/config.toml"
)

func main() {
	config := configs.NewConfig()
	_, err := toml.DecodeFile(configsPath, config)
	if err != nil {
		logrus.Error(err)
	}
}
