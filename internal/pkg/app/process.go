package app

import (
	"log"

	"github.com/Gena97/telegram_bot/internal/app/bot"
	"github.com/Gena97/telegram_bot/internal/app/config"
	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/Gena97/telegram_bot/internal/service"
)

func handleCommand(s *Service, fullEndpoint string, msg *model.Message) {
	args := service.ParseCommandArgs(msg.Text)
	botTag := "@" + config.TelegramBotTag()

	switch args[0] {
	case "/start", "/start" + botTag:
		err := bot.SendMessage(fullEndpoint, msg.ChatID, "Welcome to the bot!", "", 0, 0)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	case "/getquote", "/getquote" + botTag:
		processQuote(s, msg)
	case "/mp3":
		processMP3(s, fullEndpoint, msg)
	case "/post":
		if user, ok := s.Users[msg.FromID]; ok && user.IsAdmin {
			processPost(s, fullEndpoint, msg)
		}
	case "/cbr", "/cbr" + botTag:
		processCBR(fullEndpoint, msg.ChatID)
	case "/novichokstats", "/novichokstats" + botTag:
		processNovichokstats(fullEndpoint, msg.ChatID)
	case "/cutvideo", "/cutvideo" + botTag:
		processCutVideo(fullEndpoint, msg)
	case "/миша":
		for i := 1; i <= 10; i++ {
			err := bot.SendMessage(fullEndpoint, msg.ChatID, "@Miner228 пидарас", "", 0, 0)
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
	default:
	}
}

func handleCommandAM(s *Service, fullEndpoint string, msg *model.Message) {
	user, ok := s.Users[msg.FromID]
	if !ok {
		log.Println("Пользоваетль не найден в таблице users")
	}
	var err error
	args := service.ParseCommandArgs(msg.Text)
	botTag := "@" + config.TelegramBotTagAM()

	switch args[0] {
	case "/start", "/start" + botTag:
		if msg.ChatType == model.ChatTypePrivate || user.IsAdmin {
			err = bot.SendMessage(fullEndpoint, msg.ChatID, model.GetStartMessageAM(), "", 0, 0)
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
	case "/memes", "/memes" + botTag:
		if msg.ChatType == model.ChatTypePrivate || user.IsAdmin {
			err = bot.SendMessage(fullEndpoint, msg.ChatID, model.GetMemes(), "HTML", 0, 0)
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
	case "/lor", "/lor" + botTag:
		if msg.ChatType == model.ChatTypePrivate || user.IsAdmin {
			bot.SendMessage(fullEndpoint, msg.ChatID, model.GetLor(), "HTML", 0, 0)
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
	case "/mp3", "/mp3" + botTag:
		processMP3(s, fullEndpoint, msg)
	case "/cutvideo", "/cutvideo" + botTag:
		processCutVideo(fullEndpoint, msg)
		/*
			case "day", "week", "month", "year", "alltime", "allactive":
				active, err := db.GetActive(db, msg)
				if err != nil {
					log.Printf("Ошибка получения актива %s", err)
				} else {
					bot.SendMessage(s.TelegramBotAM.FullEndpoint, chatID, active)
				}
		*/
	default:
	}
}
