package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Gena97/telegram_bot/internal/api"
	"github.com/Gena97/telegram_bot/internal/app/bot"
	"github.com/Gena97/telegram_bot/internal/app/config"
	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/Gena97/telegram_bot/internal/pkg/scrappers"
	"github.com/Gena97/telegram_bot/internal/service"
)

func ProcessUpdates(s *Service) {
	for update := range s.TelegramBot.TelegramBotChan {
		msg := service.GetMessageFromUpdate(&update)
		if service.IsCommand(msg.Text) {
			handleCommand(s, s.TelegramBot.FullEndpoint, &msg)
		} else {
			handleMessageContent(s, s.TelegramBot.FullEndpoint, &msg)
		}
	}
}

func ProcessUpdatesAM(s *Service) {
	for update := range s.TelegramBotAM.TelegramBotChan {
		msg := service.GetMessageFromUpdate(&update)

		_, ok := s.Users[msg.FromID]
		if !ok {
			log.Println("Пользоваетль не найден в таблице users")
		}

		if service.IsCommand(msg.Text) {
			handleCommandAM(s, s.TelegramBotAM.FullEndpoint, &msg)
		} else {
			handleMessageContent(s, s.TelegramBotAM.FullEndpoint, &msg)
		}
	}
}

func handleMessageContent(s *Service, fullEndpoint string, msg *model.Message) {
	contentType := service.DetectContentType(msg.Text)
	if contentType == "" {
		return
	}
	var err error
	var mediaContentConfig model.MediaContentConfig

	switch contentType {
	case "youtube":
		if strings.Contains(msg.Text, "playlist") && !s.Users[msg.FromID].IsAdmin {
			return
		}
		mediaContentConfig, err = scrappers.DownloadVideoYoutubeV2(msg.Text)
	case "twitter":
		mediaContentConfig, err = scrappers.DownloadContentTwitter(msg.Text, s.TwitterScrapper)
	case "instagram":
		mediaContentConfig, err = scrappers.DownloadContentInstagram(msg.Text)
	default:
		return
	}
	if err != nil {
		log.Println(err)
	}
	mediaContentConfig.FullEndpoint = fullEndpoint
	mediaContentConfig.ChatID = msg.ChatID
	mediaContentConfig.Sender = msg.FromFirstName
	mediaContentConfig.ReplyToMessageID = msg.ReplyToMessageID
	bot.SendMediaContent(mediaContentConfig, msg.MessageID)
}

func processMP3(s *Service, fullEndpoint string, msg *model.Message) {
	args := service.ParseCommandArgs(msg.Text)
	if len(args) == 2 {
		if strings.Contains(args[1], "playlist") && !s.Users[msg.FromID].IsAdmin {
			return
		}

		var mediaContentConfig model.MediaContentConfig
		audioConfig, err := scrappers.DownloadAudioYoutube(args[1])
		if err != nil {
			log.Printf("Error sending message: %v", err)
			return
		}
		mediaContentConfig.FullEndpoint = fullEndpoint
		mediaContentConfig.ChatID = msg.ChatID
		mediaContentConfig.Sender = msg.FromFirstName
		mediaContentConfig.ReplyToMessageID = msg.ReplyToMessageID
		mediaContentConfig.AudioConfigs = append(mediaContentConfig.AudioConfigs, audioConfig...)
		mediaContentConfig.Link = args[1]
		mediaContentConfig.Title = "YouTube"
		bot.SendMediaContent(mediaContentConfig, msg.MessageID)
	}
}

func processQuote(s *Service, msg *model.Message) {
	quote, err := api.GetQuote()
	if err != nil {
		log.Printf("Ошибка получения цитаты.")
		return
	}
	bot.SendMessage(s.TelegramBot.FullEndpoint, msg.ChatID, quote, "", 0, 0)
}

func processPost(s *Service, fullEndpoint string, msg *model.Message) {
	fileInfo, err := bot.GetAndDownloadReplyMedia(s.TelegramBot, msg.RawMsg)
	if err != nil {
		log.Printf("Error downloading repliy media: %v", err)
		return
	}
	fmt.Println(fileInfo)

	if fileInfo.Type == "video" {
		videoHashes, err := service.ComputeVideoHashes(fileInfo.FilePath)
		if err != nil {
			log.Printf("Error getting video hashes: %v", err)
		}
		fmt.Println(len(videoHashes))

		filesToCompare, err := s.PGXMain.GetFileHashes((len(videoHashes)), "video")
		if err != nil {
			log.Printf("Error getting video hashes from DB: %v", err)
		}
		for _, toCompare := range filesToCompare {
			distance, err := service.CalculateAverageDistance(videoHashes, toCompare.Hashes)
			if err != nil {
				log.Printf("Ошибка при вычислении средней дистанции: %v", err)
			}
			if distance < 10 {
				err := bot.SendMessage(s.TelegramBot.FullEndpoint, msg.ChatID, fmt.Sprintf("Найден <a href=\"https://t.me/%s/%d\">дубликат</a>", config.TelegramBotTag(), toCompare.MessageID), "HTML", msg.ReplyToMessageID, msg.MessageID)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
				log.Printf("Найден дубликат")
				return
			}
		}
		newMsgIDs, err := bot.SendMediaContent(model.MediaContentConfig{FullEndpoint: fullEndpoint, VideosConfigs: []model.VideoConfig{{FilePath: fileInfo.FilePath}}, ChatID: model.MemniyRayChatID}, msg.MessageID)
		if err != nil {
			log.Printf("Ошибка отправки видео на канал: %v", err)
			return
		}
		err = s.PGXMain.InsertHash(service.FileHash{Hashes: videoHashes, MessageID: newMsgIDs[0], Type: "video"})
		if err != nil {
			log.Printf("Ошибка cохранения хеша фото: %v", err)
			return
		}
		err = bot.SendMessage(s.TelegramBot.FullEndpoint, msg.ChatID, "Медиа успешно отправлено в @"+model.MemniyRayTag, "HTML", msg.ReplyToMessageID, msg.MessageID)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
	if fileInfo.Type == "photo" {
		photoHash, err := service.ComputeImageHash(fileInfo.FilePath)
		if err != nil {
			log.Printf("Error getting video hashes: %v", err)
		}
		fmt.Println(photoHash)
		filesToCompare, err := s.PGXMain.GetFileHashes(1, "photo")
		if err != nil {
			log.Printf("Error getting video hashes from DB: %v", err)
		}
		for _, toCompare := range filesToCompare {
			distance, err := service.CalculateAverageDistance([]uint64{photoHash}, toCompare.Hashes)
			if err != nil {
				log.Printf("Ошибка при вычислении средней дистанции: %v", err)
			}
			if distance < 10 {
				err := bot.SendMessage(s.TelegramBot.FullEndpoint, msg.ChatID, fmt.Sprintf("Найден <a href=\"https://t.me/%s/%d\">дубликат</a>", model.MemniyRayTag, toCompare.MessageID), "HTML", msg.ReplyToMessageID, msg.MessageID)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
				log.Printf("Найден дубликат")
				return
			}
		}
		newMsgIDs, err := bot.SendMediaContent(model.MediaContentConfig{FullEndpoint: fullEndpoint, PhotosConfigs: []model.PhotoConfig{{FilePath: fileInfo.FilePath}}, ChatID: model.MemniyRayChatID}, msg.MessageID)
		if err != nil {
			log.Printf("Ошибка отправки фото на канал: %v", err)
			return
		}
		err = s.PGXMain.InsertHash(service.FileHash{Hashes: []uint64{photoHash}, MessageID: newMsgIDs[0], Type: "photo"})
		if err != nil {
			log.Printf("Ошибка cохранения хеша фото: %v", err)
			return
		}
		err = bot.SendMessage(s.TelegramBot.FullEndpoint, msg.ChatID, "Медиа успешно отправлено в @"+model.MemniyRayTag, "HTML", msg.ReplyToMessageID, msg.MessageID)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}

func processCBR(fullEndpoint string, chatID int64) {
	text, err := service.GetCurrencyRates()
	if err != nil {
		log.Printf("Error getting currency rates: %v", err)
	}
	err = bot.SendMessage(fullEndpoint, chatID, text, "", 0, 0)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func processNovichokstats(fullEndpoint string, chatID int64) {
	text, err := service.GetPubgStats()
	if err != nil {
		log.Printf("Error getting pubg stats: %v", err)
	}
	err = bot.SendMessage(fullEndpoint, chatID, text, "", 0, 0)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func processCutVideo(fullEndpoint string, msg *model.Message) {
	args := strings.Split(msg.Text, " ")
	if len(args) < 4 {
		err := bot.SendMessage(fullEndpoint, msg.ChatID, "Please specify the video URL and start and end timestamps in the format: /cutvideo <URL> <start> <end>", "", msg.ReplyToMessageID, 0)
		if err != nil {
			log.Printf("[processCutVideo] Error sending message: %v", err)
		}
		return
	}
	videoURL, startTime, endTime := args[1], args[2], args[3]

	mediaConfig, err := scrappers.DownloadVideoYoutubeV2(videoURL)
	if err != nil {
		log.Printf("[processCutVideo] Error downloading video: %s", err)
	}

	videoPath := mediaConfig.VideosConfigs[0].FilePath

	outputPath := "cut_" + filepath.Base(videoPath)
	err = service.CutVideo(videoPath, outputPath, startTime, endTime)
	if err != nil {
		log.Printf("[processCutVideo] Error cutting video: %s", err)
	}

	err = os.Remove(videoPath)
	if err != nil {
		log.Printf("[processCutVideo] Failed to delete the output file: %s", err)
	}
	mediaConfig.VideosConfigs[0].FilePath = outputPath

	mediaConfig.FullEndpoint = fullEndpoint
	mediaConfig.ChatID = msg.ChatID
	mediaConfig.Sender = msg.FromFirstName
	mediaConfig.ReplyToMessageID = msg.ReplyToMessageID

	_, err = bot.SendMediaContent(mediaConfig, msg.ReplyToMessageID)
	if err != nil {
		log.Printf("[processCutVideo] Error sending mediaConfig: %v", err)
	}
}
