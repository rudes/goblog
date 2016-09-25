package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

func getConfig() (Config, error) {
	var conf Config
	_, err := os.Stat(_staticRoot + "config.toml")
	if err != nil {
		return conf, err
	}

	_, err = toml.DecodeFile(_staticRoot+"config.toml", &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
