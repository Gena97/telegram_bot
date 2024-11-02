package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TelegramBotToken        string `yaml:"telegram_bot_token"`
	TelegramBotEndpoint     string `yaml:"telegram_bot_endpoint"` // https://api.telegram.org/bot or http://localhost:8081/bot
	TelegramServerID        string `yaml:"telegram_server_id"`
	TelegramServerHash      string `yaml:"telegram_server_hash"`
	TwitterScrapperLogin    string `yaml:"scrapper_twitter_login"`
	TwitterScrapperPassword string `yaml:"scrapper_twitter_password"`
	DBUrl                   string `yaml:"db_url"`
	LogLevel                string `yaml:"log_level"`
}

// Load загружает конфигурацию из файла config.yaml
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// Логируем уровень логирования, например
	log.Printf("Configuration loaded with log level: %s", cfg.LogLevel)

	return &cfg, nil
}
