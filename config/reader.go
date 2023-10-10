package config

import (
	"api/types"
	"os"

	"gopkg.in/yaml.v3"
)

func ConfigDB(path string) (*types.ConfDB, error) {
	var conf types.ConfDB
	file, err := os.ReadFile(path)
	if err != nil {
		return &conf, err
	}
	err = yaml.Unmarshal(file, &conf)
	return &conf, err
}

func ConfigAPI(path string) (*types.ConfAPI, error) {
	var conf types.ConfAPI
	file, err := os.ReadFile(path)
	if err != nil {
		return &conf, err
	}
	err = yaml.Unmarshal(file, &conf)
	return &conf, err
}
