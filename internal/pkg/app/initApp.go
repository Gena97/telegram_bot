package app

import (
	"log"
	"time"

	"github.com/Gena97/telegram_bot/internal/app/bot"
	"github.com/Gena97/telegram_bot/internal/app/config"
	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/Gena97/telegram_bot/internal/app/repository"
	"github.com/Gena97/telegram_bot/internal/pkg/scrappers"
	"github.com/kkdai/youtube/v2"
	twitterscraper "github.com/n0madic/twitter-scraper"
	_ "modernc.org/sqlite"
)

type Service struct {
	TelegramBot     model.TelegramBot
	TelegramBotAM   model.TelegramBot
	YoutubeClient   *youtube.Client
	TwitterScrapper *twitterscraper.Scraper
	PGXMain         repository.PGXRepository
	Users           map[int64]model.User
}

func Run() error {
	err := config.Load("../../configs/config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	s, err := initService()
	if err != nil {
		log.Fatalf("error initting service: %v", err)
	}

	go ProcessUpdates(&s)

	go ProcessUpdatesAM(&s)

	return nil
}

func initService() (Service, error) {
	var s Service
	var err error

	err = bot.InitTelegramServer()
	if err != nil {
		log.Printf("Ошибка запуска сервера Телеграм %v", err)
		return Service{}, err
	}

	s.TelegramBot, err = bot.GetTelegramBot(config.TelegramBotEndpoint(), config.TelegramBotToken())
	if err != nil {
		return Service{}, err
	}

	time.Sleep(2 * time.Second)

	s.TelegramBotAM, err = bot.GetTelegramBot(config.TelegramBotEndpoint(), config.TelegramBotTokenAM())
	if err != nil {
		return Service{}, err
	}

	s.YoutubeClient = scrappers.GetYoutubeClient()

	s.TwitterScrapper, err = scrappers.GetTwitterScrapper(config.TwitterScrapperLogin(), config.TwitterScrapperPassword())
	if err != nil {
		return Service{}, err
	}

	s.PGXMain, err = repository.InitDatabase()
	if err != nil {
		return Service{}, err
	}

	s.Users, err = s.PGXMain.GetUserInfosMap()
	if err != nil {
		return Service{}, err
	}

	return s, nil
}
