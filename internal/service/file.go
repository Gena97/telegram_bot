package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	VideoType = "video"
	ImageType = "photo"
)

type FileFromServer struct {
	Type     string
	FilePath string
}

var VideoExtensions = []string{"mp4", "avi", "mov", "mkv"}
var ImageExtensions = []string{"jpg", "jpeg", "png", "gif", "bmp"}

func SanitizeFilename(filename string) string {
	// Список запрещенных символов
	forbiddenChars := []string{"/", "\\", ":", "：", "*", "?", "\"", "<", ">", "|", " "}

	// Замена запрещенных символов на "_"
	for _, char := range forbiddenChars {
		filename = strings.ReplaceAll(filename, char, "_")
	}

	// Возвращение очищенного имени файла
	return filename
}

// downloadMedia downloads a media file from a given URL and saves it to the specified directory.
func DownloadMedia(url, directory string) (string, error) {
	// Get the media file
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download media from URL %s: %v", url, err)
	}
	defer resp.Body.Close()

	// Extract file name from the URL and set the file path

	fileName := SanitizeFilename(filepath.Base(url))

	const maxFileNameLength = 100 // Adjust as needed
	if len(fileName) > maxFileNameLength {
		fileName = fileName[:maxFileNameLength] // Keep the extension
	}

	filePath := fmt.Sprintf("%s%s", directory, fileName)

	// Create the file locally
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file %s: %v", filePath, err)
	}
	defer file.Close()

	// Write the response body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save media to file %s: %v", filePath, err)
	}

	return filePath, nil
}

func DetermineType(url string) string {
	for _, ext := range VideoExtensions {
		if strings.Contains(url, ext) {
			return VideoType
		}
	}
	for _, ext := range ImageExtensions {
		if strings.Contains(url, ext) {
			return ImageType
		}
	}
	return "unknown"
}

func GetVideoPreview(videoPath string) ([]byte, error) {
	// Создаем временный файл для миниатюры
	tempThumbPath := videoPath + ".jpg"

	// Команда для извлечения кадра с помощью ffmpeg
	cmd := exec.Command(
		"ffmpeg",
		"-i", videoPath,
		"-vframes", "1",
		"-vf", "scale='if(gt(iw/ih,1),320,-1)':'if(gt(iw/ih,1),-1,320)',pad=320:320:(320-iw)/2:(320-ih)/2:black",
		"-q:v", "2", // Устанавливаем качество сжатия (1 - лучшее, 31 - худшее)
		tempThumbPath,
	)
	cmd.Stdout = os.Stdout // Перенаправляем stdout для отладки
	cmd.Stderr = os.Stderr // Перенаправляем stderr для отладки

	// Запускаем команду ffmpeg
	if err := cmd.Run(); err != nil {
		fmt.Println("Ошибка выполнения ffmpeg:", err)
		return nil, err
	}

	// Проверка существования файла
	if _, err := os.Stat(tempThumbPath); os.IsNotExist(err) {
		fmt.Println("Файл миниатюры не найден:", err)
		return nil, err
	}

	// Открываем миниатюру для чтения
	thumbFile, err := os.Open(tempThumbPath)
	if err != nil {
		fmt.Println("Ошибка открытия файла миниатюры:", err)
		return nil, err
	}

	// Читаем данные файла миниатюры в буфер
	thumbData, err := io.ReadAll(thumbFile)
	thumbFile.Close() // Закрываем файл сразу после чтения
	if err != nil {
		fmt.Println("Ошибка чтения файла миниатюры:", err)
		return nil, err
	}

	// Удаляем временный файл миниатюры
	if err := os.Remove(tempThumbPath); err != nil {
		fmt.Println("Ошибка удаления временного файла:", err)
	}

	// Проверка размера файла
	if len(thumbData) == 0 {
		fmt.Println("Файл миниатюры пустой")
		return nil, fmt.Errorf("файл миниатюры пустой")
	}

	return thumbData, nil
}

func GenerateSavePath(mediaType string) string {
	// Путь к директории для сохранения
	dir := "../../downloads"

	// Формируем имя файла на основе типа медиа и текущего времени
	timestamp := time.Now().Format("20060102150405")
	var extension string

	if mediaType == "photo" {
		extension = ".jpg" // или .png в зависимости от ваших предпочтений
	} else if mediaType == "video" {
		extension = ".mp4"
	}

	return filepath.Join(dir, fmt.Sprintf("%s%s", timestamp, extension))
}
