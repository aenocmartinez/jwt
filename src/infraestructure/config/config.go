package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB        Database `yaml:"Database"`
	SecretKey string   `yaml:"secretKey"`
	Issuer    string   `yaml:"issuer"`
}

type Database struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
}

func Load(filename string) (*Config, error) {
	config := &Config{}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
