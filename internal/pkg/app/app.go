package app

import (
	"log"

	"github.com/Gena97/telegram_bot/internal/app/bot"
	"github.com/Gena97/telegram_bot/internal/app/config"
	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/Gena97/telegram_bot/internal/pkg/scrappers"
	"github.com/kkdai/youtube/v2"
)

type Service struct {
	TelegramBot   model.TelegramBot
	YoutubeClient *youtube.Client
}

func Run() error {
	cfg, err := config.Load("../../configs/config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	s, err := initService(cfg)
	if err != nil {
		log.Fatalf("error initting service: %v", err)
	}

	ProcessUpdates(&s)

	return nil
}

func initService(config *config.Config) (Service, error) {
	telegramBot, err := bot.GetTelegramBot(config.TelegramBotEndpoint, config.TelegramBotToken)
	if err != nil {
		return Service{}, err
	}
	youtubeClient := scrappers.GetYoutubeClient()
	s := Service{
		TelegramBot:   telegramBot,
		YoutubeClient: youtubeClient,
	}
	return s, nil
}
