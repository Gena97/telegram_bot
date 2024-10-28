package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TelegramToken string `yaml:"telegram_token"`
	BotUrl        string `yaml:"bot_url"` // https://api.telegram.org/bot or http://localhost:8081/bot
	DBUrl         string `yaml:"db_url"`
	LogLevel      string `yaml:"log_level"`
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
