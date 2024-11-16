package service

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/tidwall/gjson"
)

// DetectContentType checks the content type of the provided URL and returns it as a string.
func DetectContentType(url string) string {
	switch {
	case strings.Contains(url, "youtu"):
		return "youtube"
	case strings.Contains(url, "https://www.x.com"), strings.Contains(url, "https://x.com"), strings.Contains(url, "twitter.com"):
		return "twitter"
	case strings.Contains(url, "inst"):
		return "instagram"
	default:
		return ""
	}
}

// isCommand проверяет, является ли текст команды.
func IsCommand(text string) bool {
	return strings.HasPrefix(text, "/")
}

func ParseCommandArgs(text string) []string {
	return strings.Split(text, " ")
}

func GetMessageFromUpdate(update *gjson.Result) model.Message {
	return model.Message{
		RawMsg:           update.Get("message"),
		Text:             update.Get("message.text").String(),
		ChatID:           update.Get("message.chat.id").Int(),
		ChatType:         update.Get("message.chat.type").String(),
		MessageID:        update.Get("message.message_id").Int(),
		ReplyToMessageID: update.Get("message.reply_to_message.message_id").Int(),
		FromFirstName:    update.Get("message.from.first_name").String(),
		FromID:           update.Get("message.from.id").Int(),
	}
}

func CutVideo(inputPath, outputPath, startTime, endTime string) error {
	// Construct the ffmpeg command with arguments
	cmdArgs := []string{
		"-ss", startTime,
		"-to", endTime,
		"-i", inputPath,
		"-c", "copy",
		outputPath,
	}

	// Create the command
	cmd := exec.Command("ffmpeg", cmdArgs...)

	// Buffer to capture standard error
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// Execute the command
	err := cmd.Run()

	// Check for errors and use stderr for detailed diagnostics
	if err != nil {
		log.Printf("Error running ffmpeg: %v, stderr: %s", err, stderr.String())
		return fmt.Errorf("error running ffmpeg: %v, stderr: %s", err, stderr.String())
	}

	return nil
}
