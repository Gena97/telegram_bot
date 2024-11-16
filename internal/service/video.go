package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
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

// extractFrames extracts frames from a video every 0.2 seconds.
func ExtractFrames(videoPath string) ([]string, error) {
	tempDir, err := ioutil.TempDir("", "frames")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary directory: %w", err)
	}

	args := map[string]interface{}{
		"vf": "fps=5", // Extract frames at 5 fps (1 frame every 0.2 seconds)
	}

	err = ffmpeg_go.Input(videoPath).
		Output(filepath.Join(tempDir, "frame-%d.jpg"), args).
		OverWriteOutput().
		Run()
	if err != nil {
		return nil, fmt.Errorf("error extracting frames: %w", err)
	}

	files, err := ioutil.ReadDir(tempDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read temporary directory: %w", err)
	}

	var framePaths []string
	for _, file := range files {
		if !file.IsDir() {
			framePaths = append(framePaths, filepath.Join(tempDir, file.Name()))
		}
	}

	return framePaths, nil
}
