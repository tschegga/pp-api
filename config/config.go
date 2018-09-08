package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

var config *Config

type Config struct {
	ListeningAddr string `yaml:"listening_addr"`
	DbConnectionString string `yaml:"db_connection_string"`
	LogLevel int `yaml:"log_level"`
}

func LoadConfig(filename string) *Config {
	file_data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	config = &(Config{})

	err2 := yaml.Unmarshal(file_data, config)
	if err2 != nil {
		log.Fatal(err2)
	}
	return config
}

func GetConfig() *Config {
	if config == nil {
		log.Fatal("you have to load the config first (LoadConfig())")
	}
	return config
}
