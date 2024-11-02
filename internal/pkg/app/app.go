package app

import (
	"log"

	"github.com/Gena97/telegram_bot/internal/app/bot"
	"github.com/Gena97/telegram_bot/internal/app/config"
	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/Gena97/telegram_bot/internal/pkg/scrappers"
	"github.com/kkdai/youtube/v2"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

type Service struct {
	TelegramBot     model.TelegramBot
	YoutubeClient   *youtube.Client
	TwitterScrapper *twitterscraper.Scraper
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
	var s Service
	var err error

	err = bot.InitTelegramServer(config)
	if err != nil {
		log.Printf("Ошибка запуска сервера Телеграм %v", err)
		config.TelegramBotEndpoint = "https://api.telegram.org/bot"
	}

	s.TelegramBot, err = bot.GetTelegramBot(config.TelegramBotEndpoint, config.TelegramBotToken)
	if err != nil {
		return Service{}, err
	}

	s.YoutubeClient = scrappers.GetYoutubeClient()

	s.TwitterScrapper, err = scrappers.GetTwitterScrapper(config.TwitterScrapperLogin, config.TwitterScrapperPassword)
	if err != nil {
		return Service{}, err
	}

	return s, nil
}
