package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger     LoggerConf     `yaml:"logger"`
	GRPCServer GRPCServerConf `yaml:"grpcServer"`
}

type LoggerConf struct {
	Level string `yaml:"level"`
}

type GRPCServerConf struct {
	Address string `yaml:"address"`
}

func NewConfig(configFile string) Config {
	config := Config{}
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatal("error unmarshaling the configuration file")
	}
	return config
}
