package scrappers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/Gena97/telegram_bot/internal/service"
)

func DownloadContentInstagram(url string) (model.MediaContentConfig, error) {
	var instContent model.MediaContentConfig

	cmd := exec.Command("node", "../../instagram-direct-url-main/node_modules/instagram-url-direct/src/test.js", url)
	output, err := cmd.CombinedOutput()
	outputStr := string(output)
	log.Println(outputStr)
	if err != nil {
		return model.MediaContentConfig{}, fmt.Errorf("error executing command: %v", err)
	}
	urlRegex := regexp.MustCompile(`https?://[^\s"]+`)

	// Ищем все совпадения в строке
	matches := urlRegex.FindAllString(outputStr, -1)

	idFile := 1
	if len(matches) > 0 {
		for _, mediaURL := range matches {
			mediaType := service.DetermineType(mediaURL)
			if mediaType == service.ImageType {
				filepath, err := service.DownloadMedia(mediaURL, fmt.Sprintf("%s/%s_%d", os.TempDir(), "tempTgBot_", idFile))
				if err != nil {
					return model.MediaContentConfig{}, fmt.Errorf("error downloading file: %v", err)
				}
				idFile++
				photoConfig := model.PhotoConfig{FilePath: filepath, Title: "inst", PhotoURL: url}
				instContent.PhotosConfigs = append(instContent.PhotosConfigs, photoConfig)
			}

			if mediaType == service.VideoType {
				filepath, err := service.DownloadMedia(mediaURL, fmt.Sprintf("%s/%s_%d", os.TempDir(), "tempTgBot_", idFile))
				if err != nil {
					return model.MediaContentConfig{}, fmt.Errorf("error downloading file: %v", err)
				}
				idFile++

				duration, err := service.GetDuration(filepath)
				if err != nil {
					log.Printf("Ошибка при получении продолжительности файла: %v", err)
					return model.MediaContentConfig{}, err
				}

				videoConfig := model.VideoConfig{FilePath: filepath, Title: "inst", VideoURL: url, Duration: duration}
				instContent.VideosConfigs = append(instContent.VideosConfigs, videoConfig)
			}
		}
	}
	instContent.Title = "inst"
	instContent.Link = url

	return instContent, nil
}
