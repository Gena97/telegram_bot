package service

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

func GetDuration(filePath string) (int, error) {
	// Команда для получения информации о файле через ffmpeg
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", filePath)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0, err
	}

	// Получение вывода и преобразование его в строку
	durationStr := strings.TrimSpace(out.String())

	// Преобразование строки в float64
	durationFloat, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, err
	}

	// Преобразование float64 в int
	durationInt := int(durationFloat)

	return durationInt, nil
}
