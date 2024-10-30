package app

import (
	"log"
	"strings"

	"github.com/Gena97/telegram_bot/internal/app/bot"
	"github.com/Gena97/telegram_bot/internal/pkg/scrappers"
)

func ProcessUpdates(s *Service) {
	for update := range s.TelegramBot.TelegramBotChan { // Постоянно слушаем канал TelegramBotChan
		text := update.Get("message.text").String()
		chatID := update.Get("message.chat.id").Int()
		replyToMessageID := update.Get("message.reply_to_message.message_id").Int()
		firstName := update.Get("message.from.first_name").String()

		switch text {
		case "/start":
			err := bot.SendMessage(s.TelegramBot.FullEndpoint, chatID, "Welcome to the bot!")
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		default:
			if strings.Contains(text, "youtu") {
				args := strings.Split(text, " ")
				if len(args) == 1 {
					VideoConfig, err := scrappers.DownloadVideoYoutube(text, s.YoutubeClient)
					if err != nil {
						log.Println(err)
						continue
					}
					VideoConfig.FullEndpoint = s.TelegramBot.FullEndpoint
					VideoConfig.ChatID = chatID
					VideoConfig.ReplyToMessageID = replyToMessageID
					VideoConfig.VideoURL = text
					VideoConfig.Sender = firstName
					bot.SendVideoAndDeleteFile(VideoConfig)
				}
			}
		}
	}
}
