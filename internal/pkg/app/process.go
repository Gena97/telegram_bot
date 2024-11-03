package app

import (
	"log"

	"github.com/Gena97/telegram_bot/internal/app/bot"
	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/Gena97/telegram_bot/internal/pkg/scrappers"
	"github.com/Gena97/telegram_bot/internal/service"
)

func ProcessUpdates(s *Service) {
	for update := range s.TelegramBot.TelegramBotChan {
		text := update.Get("message.text").String()
		chatID := update.Get("message.chat.id").Int()
		replyToMessageID := update.Get("message.reply_to_message.message_id").Int()
		firstName := update.Get("message.from.first_name").String()
		messageID := update.Get("message.message_id").Int()

		if service.IsCommand(text) {
			handleCommand(s, text, s.TelegramBot.FullEndpoint, firstName, chatID, replyToMessageID, messageID)
		} else {
			handleMessageContent(s, text, chatID, replyToMessageID, firstName, messageID)
		}
	}
}

func ProcessUpdatesAM(s *Service) {
	for update := range s.TelegramBotAM.TelegramBotChan {
		text := update.Get("message.text").String()
		chatID := update.Get("message.chat.id").Int()
		replyToMessageID := update.Get("message.reply_to_message.message_id").Int()
		firstName := update.Get("message.from.first_name").String()
		messageID := update.Get("message.message_id").Int()

		if service.IsCommand(text) {
			handleCommandAM(s, text, s.TelegramBotAM.FullEndpoint, firstName, chatID, replyToMessageID, messageID)
		} else {
			handleMessageContent(s, text, chatID, replyToMessageID, firstName, messageID)
		}
	}
}

func handleCommand(s *Service, text, fullEndpoint, firstName string, chatID, replyToMessageID, messageID int64) {

	args := service.ParseCommandArgs(text)

	switch args[0] {
	case "/start":
		err := bot.SendMessage(s.TelegramBot.FullEndpoint, chatID, "Welcome to the bot!", "")
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	case "/mp3":
		var mediaContentConfig model.MediaContentConfig
		if len(args) == 2 {
			audioConfig, err := scrappers.DownloadAudioYoutube(args[1])
			if err != nil {
				log.Printf("Error sending message: %v", err)
				return
			}
			setMediaContentFields(&mediaContentConfig, fullEndpoint, chatID, firstName, replyToMessageID)
			mediaContentConfig.AudioConfigs = append(mediaContentConfig.AudioConfigs, audioConfig)
			mediaContentConfig.Link = args[1]
			mediaContentConfig.Title = "YouTube"
			bot.SendMediaContent(mediaContentConfig, messageID)
		}
	default:
	}
}

func handleMessageContent(s *Service, text string, chatID, replyToMessageID int64, firstName string, messageID int64) {
	contentType := service.DetectContentType(text)
	if contentType != "" {
		handleContentType(s, contentType, text, s.TelegramBot.FullEndpoint, firstName, chatID, replyToMessageID, messageID)
	}
}

func handleCommandAM(s *Service, text, fullEndpoint, firstName string, chatID, replyToMessageID, messageID int64) {

	args := service.ParseCommandArgs(text)

	switch text {
	case "/start":
		err := bot.SendMessage(s.TelegramBotAM.FullEndpoint, chatID, "Ты все сказал? А теперь мне слушай!\n\nЧитай лор чата, гандон - /lor\n\nПрочитал? Теперь мемы - /memes", "")
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	case "/memes":
		bot.SendMessage(s.TelegramBotAM.FullEndpoint, chatID, model.GetMemes(), "HTML")
	case "/lor":
		bot.SendMessage(s.TelegramBotAM.FullEndpoint, chatID, model.GetLor(), "HTML")
	case "/mp3":
		var mediaContentConfig model.MediaContentConfig
		if len(args) == 2 {
			audioConfig, err := scrappers.DownloadAudioYoutube(args[1])
			if err != nil {
				log.Printf("Error sending message: %v", err)
				return
			}
			setMediaContentFields(&mediaContentConfig, fullEndpoint, chatID, firstName, replyToMessageID)
			mediaContentConfig.AudioConfigs = append(mediaContentConfig.AudioConfigs, audioConfig)
			mediaContentConfig.Link = args[1]
			mediaContentConfig.Title = "YouTube"
			bot.SendMediaContent(mediaContentConfig, messageID)
		}
	default:
	}
}

// handleContentType processes content based on its type and sends a response.
func handleContentType(s *Service, contentType, text, fullEndpoint, firstName string, chatID, replyToMessageID, messageID int64) {
	switch contentType {
	case "youtube":
		mediaContentConfig, err := scrappers.DownloadVideoYoutubeV2(text)
		if err != nil {
			log.Println(err)
			return
		}
		setMediaContentFields(&mediaContentConfig, fullEndpoint, chatID, firstName, replyToMessageID)
		bot.SendMediaContent(mediaContentConfig, messageID)
	case "twitter":
		mediaContentConfig, err := scrappers.DownloadContentTwitter(text, s.TwitterScrapper)
		if err != nil {
			log.Println(err)
		}
		setMediaContentFields(&mediaContentConfig, fullEndpoint, chatID, firstName, replyToMessageID)
		bot.SendMediaContent(mediaContentConfig, messageID)
	case "instagram":
		mediaContentConfig, err := scrappers.DownloadContentInstagram(text)
		if err != nil {
			log.Println(err)
		}
		setMediaContentFields(&mediaContentConfig, fullEndpoint, chatID, firstName, replyToMessageID)
		bot.SendMediaContent(mediaContentConfig, messageID)
	}
}

// setMediaContentFields populates common fields in the media content configuration.
func setMediaContentFields(mediaContentConfig *model.MediaContentConfig, fullEndpoint string, chatID int64, sender string, replyToMessageID int64) {
	mediaContentConfig.FullEndpoint = fullEndpoint
	mediaContentConfig.ChatID = chatID
	mediaContentConfig.Sender = sender
	mediaContentConfig.ReplyToMessageID = replyToMessageID
}
