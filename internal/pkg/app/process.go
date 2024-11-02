package app

import (
	"log"
	"strings"

	"github.com/Gena97/telegram_bot/internal/app/bot"
	"github.com/Gena97/telegram_bot/internal/pkg/scrappers"
	"github.com/Gena97/telegram_bot/internal/service"
)

func ProcessUpdates(s *Service) {
	for update := range s.TelegramBot.TelegramBotChan {
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
			args := strings.Split(text, " ")
			if len(args) == 1 {
				contentType := service.DetectContentType(text)
				if contentType != "" {
					handleContentType(contentType, text, s, chatID, replyToMessageID, firstName)
				}
			}
		}
	}
}

// handleContentType processes content based on its type and sends a response.
func handleContentType(contentType, text string, s *Service, chatID, replyToMessageID int64, firstName string) {
	switch contentType {
	case "youtube":
		videoConfig, err := scrappers.DownloadVideoYoutube(text, s.YoutubeClient)
		if err != nil {
			log.Println(err)
			return
		}
		videoConfig.FullEndpoint = s.TelegramBot.FullEndpoint
		videoConfig.ChatID = chatID
		videoConfig.ReplyToMessageID = replyToMessageID
		videoConfig.VideoURL = text
		videoConfig.Sender = firstName
		bot.SendVideoAndDeleteFile(videoConfig)
	case "twitter":
		mediaContentConfig, err := scrappers.DownloadContentTwitter(text, s.TwitterScrapper)
		if err != nil {
			log.Println(err)
		}
		mediaContentConfig.FullEndpoint = s.TelegramBot.FullEndpoint
		mediaContentConfig.ChatID = chatID
		mediaContentConfig.Sender = firstName
		mediaContentConfig.ReplyToMessageID = replyToMessageID
		bot.SendMediaContent(mediaContentConfig)
	case "instagram":
		mediaContentConfig, err := scrappers.DownloadContentInstagram(text)
		if err != nil {
			log.Println(err)
		}
		mediaContentConfig.FullEndpoint = s.TelegramBot.FullEndpoint
		mediaContentConfig.ChatID = chatID
		mediaContentConfig.Sender = firstName
		mediaContentConfig.ReplyToMessageID = replyToMessageID
		bot.SendMediaContent(mediaContentConfig)
	}

}
