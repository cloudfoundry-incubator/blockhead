package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
	Port     uint16 `json:"port"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("Error opening config file: %v", err.Error())
	}

	err = json.Unmarshal(bytes, config)
	if err != nil {
		return nil, fmt.Errorf("Error parsing config file: %v", err.Error())
	}

	if config.Port == 0 {
		config.Port = 3333
	}

	return config, nil
}
