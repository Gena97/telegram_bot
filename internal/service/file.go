package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	VideoType = "video"
	ImageType = "photo"
)

var VideoExtensions = []string{"mp4", "avi", "mov", "mkv"}
var ImageExtensions = []string{"jpg", "jpeg", "png", "gif", "bmp"}

func SanitizeFilename(filename string) string {
	// Список запрещенных символов
	forbiddenChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|", " "}

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
