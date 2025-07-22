package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"path/filepath"
)

type Config struct {
	ServerConfig
}

type ServerConfig struct {
	TrustedProxies []string `json:"trusted_proxies"`
}

var (
	AppConfig *Config
	once sync.Once
)

func InitConfig(optionalPath ...string) {
	var path string
	once.Do(func() {
		AppConfig = &Config{}
	})
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
	if err := decoder.Decode(&AppConfig.ServerConfig); err != nil {
		panic(fmt.Sprintf("Implement Me: %v\n", err))
	}
}

func DatabaseConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host,
		user,
		password,
		dbname,
		port,
		sslmode,
		timezone,
	)
}
