package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

func ConfigDB(path string) (*ConfDB, error) {
	var conf ConfDB
	file, err := os.ReadFile(path)
	if err != nil {
		return &conf, err
	}
	err = yaml.Unmarshal(file, &conf)
	return &conf, err
}

func ConfigAPI(path string) (*ConfAPI, error) {
	var conf ConfAPI
	file, err := os.ReadFile(path)
	if err != nil {
		return &conf, err
	}
	err = yaml.Unmarshal(file, &conf)
	return &conf, err
}