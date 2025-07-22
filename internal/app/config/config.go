package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	TrustedProxies []string `json:"trusted_proxies"`
}

var AppConfig *Config

func InitConfig(optionalPath ...string) {
	var path string
	switch len(optionalPath) {
	case 0:
		path = filepath.Join("configs", "server.json")
	case 1:
		path = optionalPath[0]
	default:
		panic("Error: only provide one configuration file path to function.")
	}

	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Implement Me: %v\n", err))
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	AppConfig = &Config{}
	if err := decoder.Decode(AppConfig); err != nil {
		panic(fmt.Sprintf("Implement Me: %v\n", err))
	}
}
