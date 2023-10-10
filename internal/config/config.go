package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database string `yaml:"dbname"`
	Username string `yaml:"dbuser"`
	Password string `yaml:"dbpass"`
	Host     string `yaml:"dbhost"`
	Port     string `yaml:"dbport"`
}

func NewConfig(path string) (*Config, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var config Config
	return &config, yaml.NewDecoder(f).Decode(&config)
}
