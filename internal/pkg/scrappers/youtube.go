package scrappers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/Gena97/telegram_bot/internal/service"
	"github.com/kkdai/youtube/v2"
)

func GetYoutubeClient() *youtube.Client {
	client := &youtube.Client{}
	return client
}

func DownloadVideoYoutube(videoURL string, client *youtube.Client) (model.VideoConfig, error) {
	videoID := extractVideoIDFromURLYoutube(videoURL)

	video, err := client.GetVideo(videoID)
	if err != nil {
		return model.VideoConfig{}, err
	}

	// Filter formats based on desired quality
	var selectedFormat youtube.Format
	formats := video.Formats.WithAudioChannels()
	for _, format := range formats {
		// Example: Check if format's quality matches the desired quality
		if format.QualityLabel == "720p" {
			selectedFormat = format
			break
		}
	}

	if (youtube.Format{} == selectedFormat) {
		selectedFormat = formats[0]
	}

	stream, _, err := client.GetStream(video, &selectedFormat)
	if err != nil {
		return model.VideoConfig{}, err
	}

	// Загрузка видео во временное хранилище
	tempFile, err := os.Create(os.TempDir() + service.SanitizeFilename(video.Title) + ".mp4")
	if err != nil {
		return model.VideoConfig{}, err
	}
	defer tempFile.Close()

	// Запись видео во временный файл
	_, err = io.Copy(tempFile, stream)
	if err != nil {
		tempFile.Close() // Make sure to close the file on error
		return model.VideoConfig{}, err
	}

	stream.Close()
	tempFile.Close()

	videoPath := tempFile.Name()

	duration, err := service.GetDuration(videoPath)
	if err != nil {
		log.Printf("Ошибка при получении продолжительности файла: %v", err)
		return model.VideoConfig{}, err
	}

	timing, err := extractTimeFromURL(videoURL)
	if err != nil {
		log.Printf("Тайминг у видео не найден")
	}

	return model.VideoConfig{FilePath: videoPath, Title: video.Title, Duration: duration, Timing: timing}, nil
}

func extractVideoIDFromURLYoutube(url string) string {
	// Проверяем, содержит ли URL часть "watch?v="
	if strings.Contains(url, "watch?v=") {
		// Извлекаем ID видео после "watch?v="
		parts := strings.Split(url, "watch?v=")
		lastPart := parts[1]
		// Обрезаем все после "&" если есть
		if idx := strings.Index(lastPart, "&"); idx != -1 {
			return lastPart[:idx]
		}
		return lastPart
	}

	// Обработка для сокращенных URL типа "youtu.be"
	if strings.Contains(url, "youtu.be") {
		parts := strings.Split(url, "/")
		lastPart := parts[len(parts)-1]
		// Обрезаем параметры запроса, если они есть
		if idx := strings.Index(lastPart, "?"); idx != -1 {
			return lastPart[:idx]
		}
		return lastPart
	}

	// Обработка других типов URL
	parts := strings.Split(url, "/")
	lastPart := parts[len(parts)-1]
	// Обрезаем параметры запроса, если они есть
	if idx := strings.Index(lastPart, "?"); idx != -1 {
		return lastPart[:idx]
	}

	return lastPart
}

func extractTimeFromURL(videoURL string) (string, error) {
	u, err := url.Parse(videoURL)
	if err != nil {
		return "", err
	}

	queryParams := u.Query()
	timing := queryParams.Get("t")
	if timing == "" {
		// Если параметр времени отсутствует, возвращаем 00:00:00.
		return "", fmt.Errorf("тайминг отсутствует")
	}

	seconds, err := strconv.Atoi(timing)
	if err != nil {
		return "", err
	}

	// Преобразование секунд в форматированное время.
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds), nil
}

// DownloadAudio загружает аудио из видео по URL и сохраняет его в формате MP3
func DownloadAudioYoutube(url string) (model.AudioConfig, error) {
	var config model.AudioConfig

	// Директория для сохранения аудио файлов
	outputDir := "../../utilities/yt-dlp/"
	outputTemplate := outputDir + "%(title)s.%(ext)s"

	// Команда для загрузки только аудиодорожки и конвертации в MP3
	cmd := exec.Command("../../utilities/yt-dlp/yt-dlp_x86.exe", "-f", "bestaudio", "--extract-audio", "--audio-format", "mp3", "-o", outputTemplate, url)

	// Выполняем команду
	if err := cmd.Run(); err != nil {
		return model.AudioConfig{}, fmt.Errorf("error downloading audio: %w", err)
	}

	// Читаем содержимое директории для поиска самого нового mp3 файла
	files, err := ioutil.ReadDir(outputDir)
	if err != nil {
		return model.AudioConfig{}, fmt.Errorf("error reading directory: %w", err)
	}

	// Ищем самый новый файл с расширением .mp3
	var newestFile os.FileInfo
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".mp3") {
			if newestFile == nil || file.ModTime().After(newestFile.ModTime()) {
				newestFile = file
			}
		}
	}

	// Проверяем, найден ли файл
	if newestFile == nil {
		return model.AudioConfig{}, fmt.Errorf("no mp3 file found in directory")
	}

	// Устанавливаем путь к найденному файлу в конфигурации
	config.FilePath = filepath.Join(outputDir, newestFile.Name())

	duration, err := service.GetDuration(config.FilePath)
	if err != nil {
		log.Printf("Ошибка при получении продолжительности файла: %v", err)
	} else {
		config.Duration = duration
	}

	return config, nil
}

// DownloadVideo загружает видео по URL и сохраняет его в формате MP4 с максимальным качеством 720p
func DownloadVideoYoutubeV2(url string) (model.MediaContentConfig, error) {
	var mediaContentConfig model.MediaContentConfig
	var videoConfig model.VideoConfig
	// Директория для сохранения видео файлов
	outputDir := "../../utilities/yt-dlp/"
	outputTemplate := outputDir + "%(title)s.%(ext)s"

	// Команда для загрузки видео с ограничением качества 720p и сохранением в формате MP4
	cmd := exec.Command("../../utilities/yt-dlp/yt-dlp_x86.exe", "-f", "bestvideo[height<=720][ext=mp4]+bestaudio[ext=m4a]/best[height<=720][ext=mp4]", "-o", outputTemplate, url)

	// Выполняем команду
	if err := cmd.Run(); err != nil {
		return model.MediaContentConfig{}, fmt.Errorf("error downloading video: %w", err)
	}

	// Читаем содержимое директории для поиска самого нового mp4 файла
	files, err := ioutil.ReadDir(outputDir)
	if err != nil {
		return model.MediaContentConfig{}, fmt.Errorf("error reading directory: %w", err)
	}

	// Ищем самый новый файл с расширением .mp4
	var newestFile os.FileInfo
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".mp4") {
			if newestFile == nil || file.ModTime().After(newestFile.ModTime()) {
				newestFile = file
			}
		}
	}

	// Проверяем, найден ли файл
	if newestFile == nil {
		return model.MediaContentConfig{}, fmt.Errorf("no mp4 file found in directory")
	}

	timing, err := extractTimeFromURL(url)
	if err != nil {
		log.Printf("Тайминг у видео не найден")
	}
	videoConfig.FilePath = filepath.Join(outputDir, newestFile.Name())
	duration, err := service.GetDuration(videoConfig.FilePath)
	if err != nil {
		log.Printf("Ошибка при получении продолжительности файла: %v", err)
	} else {
		videoConfig.Duration = duration
	}

	// Устанавливаем путь к найденному файлу в конфигурации

	videoConfig.Timing = timing

	mediaContentConfig.VideosConfigs = append(mediaContentConfig.VideosConfigs, videoConfig)
	mediaContentConfig.Link = url
	mediaContentConfig.Title = newestFile.Name()

	return mediaContentConfig, nil
}
