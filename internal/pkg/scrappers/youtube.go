package scrappers

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/kkdai/youtube/v2"
	"golang.org/x/exp/rand"

	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/Gena97/telegram_bot/internal/service"
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
		log.Printf("Тайминг в ссылке ютуб видео не найден")
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

	// Проверка формата `число` (в секундах)
	if seconds, err := strconv.Atoi(timing); err == nil {
		hours := seconds / 3600
		minutes := (seconds % 3600) / 60
		seconds = seconds % 60
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds), nil
	}

	// Проверка формата `XhYmZs`, `XmYs` или только `Xs`
	re := regexp.MustCompile(`(?:(\d+)h)?(?:(\d+)m)?(?:(\d+)s)?`)
	matches := re.FindStringSubmatch(timing)
	if matches == nil {
		return "", fmt.Errorf("неизвестный формат времени: %s", timing)
	}

	// Преобразование часов, минут и секунд из строки в целое число (если поле пустое, присваиваем 0)
	hours, _ := strconv.Atoi(matches[1])
	minutes, _ := strconv.Atoi(matches[2])
	seconds, _ := strconv.Atoi(matches[3])

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds), nil
}

// DownloadAudio загружает аудио из видео по URL и сохраняет его в формате MP3
func DownloadAudioYoutube(url string) ([]model.AudioConfig, error) {
	var configs []model.AudioConfig

	// Генерируем случайный префикс для имени папки
	rand.Seed(uint64(time.Now().UnixNano()))
	randNum := rand.Intn(100000)
	strRanNum := strconv.Itoa(randNum)

	outputDir := "../../utilities/yt-dlp/"
	// Шаблон для сохранения файла
	outputTemplate := filepath.Join(outputDir, "%(title)s"+strRanNum)

	// Команда для загрузки только аудиодорожки и конвертации в MP3
	cmd := exec.Command("../../utilities/yt-dlp/yt-dlp_x86.exe", "-f", "bestaudio", "--extract-audio", "--audio-format", "mp3", "-o", outputTemplate, url)

	// Выполняем команду
	if err := cmd.Run(); err != nil {
		return []model.AudioConfig{}, fmt.Errorf("error downloading audio: %w", err)
	}

	// Читаем содержимое директории для поиска самого нового mp3 файла
	files, err := os.ReadDir(outputDir)
	if err != nil {
		return []model.AudioConfig{}, fmt.Errorf("error reading directory: %w", err)
	}

	// Ищем самый новый файл с расширением .mp3
	var newestFiles []fs.DirEntry
	for _, file := range files {
		if strings.Contains(file.Name(), strRanNum) {
			newestFiles = append(newestFiles, file)
		}
	}

	// Проверяем, найден ли файл
	if newestFiles == nil {
		return []model.AudioConfig{}, fmt.Errorf("no mp3 file found in directory")
	}

	// Сортируем по времени создания
	sort.Slice(newestFiles, func(i, j int) bool {
		info1, err1 := files[i].Info()
		info2, err2 := files[j].Info()
		if err1 != nil || err2 != nil {
			panic("Failed to retrieve file info")
		}
		return info1.ModTime().Before(info2.ModTime())
	})

	for _, newestFile := range newestFiles {
		var config model.AudioConfig
		// Устанавливаем путь к найденному файлу в конфигурации
		config.FilePath = filepath.Join(outputDir, newestFile.Name())
		duration, err := service.GetDuration(config.FilePath)
		if err != nil {
			log.Printf("Ошибка при получении продолжительности файла: %v", err)
		} else {
			config.Duration = duration
		}
		config.Title = newestFile.Name()[:len(newestFile.Name())-(5+len(strRanNum))]
		configs = append(configs, config)
	}

	return configs, nil
}

// DownloadVideo загружает видео по URL и сохраняет его в формате MP4 с максимальным качеством 720p
func DownloadVideoYoutubeV2(url string) (model.MediaContentConfig, error) {
	var mediaContentConfig model.MediaContentConfig
	var videoConfig model.VideoConfig

	// Директория для сохранения видео файлов

	// Генерируем случайный префикс для имени папки
	rand.Seed(uint64(time.Now().UnixNano()))
	randNum := rand.Intn(100000)
	strRanNum := strconv.Itoa(randNum)

	outputDir := "../../utilities/yt-dlp/"

	// Шаблон для сохранения файла
	outputTemplate := filepath.Join(outputDir, "%(title)s.%(ext)s"+strRanNum)

	// Команда для загрузки видео с ограничением качества 720p и сохранением в формате MP4
	cmd := exec.Command(
		"../../utilities/yt-dlp/yt-dlp_x86.exe",
		"-f", "bestvideo[height<=720][ext=mp4]+bestaudio[ext=m4a]/best[height<=720][ext=mp4]",
		"-o", outputTemplate, // Шаблон имени файла
		url,
	)

	// Выполняем команду
	if err := cmd.Run(); err != nil {
		return model.MediaContentConfig{}, fmt.Errorf("error downloading video: %w", err)
	}

	// Читаем содержимое директории для поиска самого нового mp4 файла
	files, err := os.ReadDir(outputDir)
	if err != nil {
		return model.MediaContentConfig{}, fmt.Errorf("error reading directory: %w", err)
	}

	var newFile fs.DirEntry

	for _, file := range files {
		if strings.Contains(file.Name(), strRanNum) {
			newFile = file
			break
		}
	}

	// Применяем sanitize к имени файла
	sanitizedFileName := service.SanitizeFilename(newFile.Name())

	// Получаем абсолютный путь для нового имени файла
	oldFilePath := filepath.Join(outputDir, newFile.Name())
	newFilePath := filepath.Join(outputDir, sanitizedFileName)

	// Переименовываем файл
	if err := os.Rename(oldFilePath, newFilePath); err != nil {
		return model.MediaContentConfig{}, fmt.Errorf("error renaming file: %w", err)
	}

	timing, err := extractTimeFromURL(url)
	if err != nil {
		log.Printf("Тайминг у видео не найден")
	}

	absoluteFilePath, err := filepath.Abs(filepath.Join(outputDir, sanitizedFileName))
	if err != nil {
		return model.MediaContentConfig{}, fmt.Errorf("error getting absolute path: %w", err)
	}

	videoConfig.FilePath = absoluteFilePath

	duration, err := service.GetDuration(videoConfig.FilePath)
	if err != nil {
		log.Printf("Ошибка при получении продолжительности файла: %v", err)
	} else {
		videoConfig.Duration = duration
	}

	// Устанавливаем путь к найденному файлу в конфигурации

	videoConfig.Timing = timing
	videoConfig.Thumbnail, err = service.GetVideoPreview(videoConfig.FilePath)
	if err != nil {
		log.Printf("ошибка получения превью видео:%v", err)
	}
	mediaContentConfig.VideosConfigs = append(mediaContentConfig.VideosConfigs, videoConfig)
	mediaContentConfig.Link = url
	mediaContentConfig.Title = newFile.Name()[:len(newFile.Name())-(8+len(strRanNum))]

	return mediaContentConfig, nil
}
