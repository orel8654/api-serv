package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// ОПИШЕМ КОНИФИГ

// создадим конструктор к тому как инициализировать этот конфиг

type Config struct {
	DB    string
	Redis string
	// ...
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
