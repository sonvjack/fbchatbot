package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AccessToken string `yaml:"accessToken"`
	VerifyToken string `yaml:"verifyToken"`
	PageID      string `yaml:"pageId"`
	LogPath     string `yaml:"logPath"`
	Port        int    `yaml:"port"`
}

func NewConfig(configF string) *Config {
	dataBytes, err := os.ReadFile(configF)
	if err != nil {
		fmt.Println("Read path:error", configF, err)
		return nil
	}
	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		fmt.Println("Parse yaml format failed:", err)
		return nil
	}
	return &config
}
