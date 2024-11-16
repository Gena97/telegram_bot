package service

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/corona10/goimagehash"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type FileHash struct {
	Hashes    []uint64
	MessageID int64
	Type      string
}

// extractFrames extracts frames from a video every 0.2 seconds.
func extractFrames(videoPath string) ([]string, error) {
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

// computeVideoHashes extracts frames from a video and computes perceptual hashes.
func ComputeVideoHashes(videoPath string) ([]uint64, error) {
	framePaths, err := extractFrames(videoPath)
	if err != nil {
		log.Printf("Error extracting frames from video %s: %v", videoPath, err)
		return nil, err
	}

	time.Sleep(time.Second)

	var hashStrs []uint64
	for _, framePath := range framePaths {
		frame, err := os.Open(framePath)
		if err != nil {
			log.Printf("Error opening frame: %s, error: %v", framePath, err)
			continue
		}

		img, err := jpeg.Decode(frame)
		frame.Close()
		if err != nil {
			log.Printf("Error decoding image: %v", err)
			continue
		}

		hash, err := goimagehash.PerceptionHash(img)
		if err != nil {
			log.Printf("Error creating hash for frame: %v", err)
			continue
		}

		hashStrs = append(hashStrs, hash.GetHash())
	}

	return hashStrs, nil
}

// computeImageHash computes the perceptual hash of an image.
func ComputeImageHash(path string) (uint64, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening file %s: %v", path, err)
		return 0, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Printf("Error decoding image %s: %v", path, err)
		return 0, err
	}

	hash, err := goimagehash.PerceptionHash(img)
	if err != nil {
		log.Printf("Error creating pHash for image %s: %v", path, err)
		return 0, err
	}

	return hash.GetHash(), nil
}

// Вычисляет среднюю дистанцию между двумя наборами хешей.
func CalculateAverageDistance(hashes1, hashes2 []uint64) (int, error) {
	if len(hashes1) == 0 || len(hashes2) == 0 || len(hashes1) != len(hashes2) {
		return 0, fmt.Errorf("неверное количество хешей")
	}

	totalDistance := 0
	for i, hashStr1 := range hashes1 {
		hash1 := goimagehash.NewImageHash(hashStr1, goimagehash.PHash)
		hash2 := goimagehash.NewImageHash(hashes2[i], goimagehash.PHash)
		distance, err := hash1.Distance(hash2)
		if err != nil {
			return 0, err
		}
		totalDistance += distance
	}

	averageDistance := totalDistance / len(hashes1)
	return averageDistance, nil
}
