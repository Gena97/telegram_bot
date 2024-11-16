package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var config appConfig

type appConfig struct {
	TelegramBotToken        string `yaml:"telegram_bot_token"`
	TelegramBotTokenAM      string `yaml:"telegram_bot_token_am"`
	TelegramBotTag          string `yaml:"telegram_bot_tag"`
	TelegramBotTagAM        string `yaml:"telegram_bot_tag_am"`
	TelegramBotEndpoint     string `yaml:"telegram_bot_endpoint"`
	TelegramServerID        string `yaml:"telegram_server_id"`
	TelegramServerHash      string `yaml:"telegram_server_hash"`
	TwitterScrapperLogin    string `yaml:"scrapper_twitter_login"`
	TwitterScrapperPassword string `yaml:"scrapper_twitter_password"`
	ApiPubgAccID            string `yaml:"api_pubg_accid"`
	ApiPubgSeasonID         string `yaml:"api_pubg_seasonId"`
	ApiPubgKey              string `yaml:"api_pubg_key"`
}

func Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	log.Printf("Configuration loaded")

	return nil
}

func TelegramBotToken() string {
	return config.TelegramBotToken
}

func TelegramBotTokenAM() string {
	return config.TelegramBotTokenAM
}

func TelegramBotTag() string {
	return config.TelegramBotTag
}

func TelegramBotTagAM() string {
	return config.TelegramBotTagAM
}

func TelegramBotEndpoint() string {
	return config.TelegramBotEndpoint
}

func TelegramServerID() string {
	return config.TelegramServerID
}

func TelegramServerHash() string {
	return config.TelegramServerHash
}

func TwitterScrapperLogin() string {
	return config.TwitterScrapperLogin
}

func TwitterScrapperPassword() string {
	return config.TwitterScrapperPassword
}

func ApiPubgAccID() string {
	return config.ApiPubgAccID
}

func ApiPubgSeasonID() string {
	return config.ApiPubgSeasonID
}

func ApiPubgKey() string {
	return config.ApiPubgKey
}
