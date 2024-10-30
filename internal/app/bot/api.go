package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/tidwall/gjson"
)

// getUpdates получает обновления от Telegram API
func GetUpdates(apiURL string, offset int, allowedUpdates []string) ([]gjson.Result, error) {
	// Формируем параметры запроса
	params := url.Values{}
	params.Add("offset", fmt.Sprintf("%d", offset))

	if len(allowedUpdates) > 0 {
		// Сериализуем `allowed_updates` как JSON-строку
		allowedUpdatesJSON, err := json.Marshal(allowedUpdates)
		if err != nil {
			return nil, fmt.Errorf("error marshaling allowed_updates: %v", err)
		}
		params.Add("allowed_updates", string(allowedUpdatesJSON))
	}

	// Формируем конечный URL с параметрами
	endpoint := fmt.Sprintf("%s/getUpdates?%s", apiURL, params.Encode())
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Читаем и обрабатываем ответ
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Извлекаем обновления из ответа
	updates := gjson.Get(string(b), "result").Array()

	return updates, nil
}

// sendMessage отправляет сообщение через Telegram API
func SendMessage(apiURL string, chatID int64, text string) error {
	endpoint := fmt.Sprintf("%s/sendMessage", apiURL)

	data := url.Values{}
	data.Set("chat_id", strconv.FormatInt(chatID, 10))
	data.Set("text", text)

	resp, err := http.PostForm(endpoint, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		Ok bool `json:"ok"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if !result.Ok {
		return fmt.Errorf("failed to send message")
	}

	return nil
}

// SendVideoAndDeleteFile отправляет видео через Telegram API с дополнительными параметрами
func SendVideoAndDeleteFile(videoConfig model.VideoConfig) error {
	endpoint := fmt.Sprintf("%s/sendVideo", videoConfig.FullEndpoint)

	// Open the video file
	file, err := os.Open(videoConfig.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a multipart form request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the video file to the form
	part, err := writer.CreateFormFile("video", filepath.Base(videoConfig.FilePath))
	if err != nil {
		return err
	}
	if _, err := io.Copy(part, file); err != nil {
		return err
	}

	// Add required fields
	writer.WriteField("chat_id", strconv.FormatInt(videoConfig.ChatID, 10))

	// Add optional fields if they are provided
	if videoConfig.ReplyToMessageID != 0 {
		writer.WriteField("reply_to_message_id", strconv.Itoa(int(videoConfig.ReplyToMessageID)))
	}
	caption := fmt.Sprintf("<a href='%s'>%s</a>", videoConfig.VideoURL, videoConfig.Title)

	if videoConfig.Timing != "" {
		caption += fmt.Sprintf("\n\n%s", videoConfig.Timing)
	}

	if videoConfig.Sender != "" {
		caption += "\n\nОтправил " + videoConfig.Sender
	}
	if videoConfig.Title != "" {
		writer.WriteField("caption", caption)
	}

	if videoConfig.Duration != 0 {
		writer.WriteField("duration", strconv.Itoa(videoConfig.Duration))
	}

	writer.WriteField("parse_mode", "HTML")
	writer.WriteField("supports_streaming", "true")

	// Finalize the form
	writer.Close()

	// Create and send the request
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		Ok bool `json:"ok"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if !result.Ok {
		return fmt.Errorf("failed to send video")
	}

	// Delete the file after sending
	defer os.Remove(videoConfig.FilePath)

	return nil
}
