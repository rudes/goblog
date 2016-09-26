package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

func getConfig(cfgFile string) (Config, error) {
	var conf Config
	_, err := os.Stat(cfgFile)
	if err != nil {
		return conf, err
	}

	toml.DecodeFile(cfgFile, &conf)
	return conf, nil
}
