package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"

	"github.com/Gena97/telegram_bot/internal/app/model"
)

// getUpdates получает обновления от Telegram API
func GetUpdates(apiURL string, offset int64, allowedUpdates []string) ([]gjson.Result, error) {
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
func SendMessage(apiURL string, chatID int64, text, parse_mode string) error {
	endpoint := fmt.Sprintf("%s/sendMessage", apiURL)

	data := url.Values{}
	data.Set("chat_id", strconv.FormatInt(chatID, 10))
	data.Set("text", text)
	if parse_mode != "" {
		data.Set("parse_mode", "HTML")
	}

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

// SendVideoAndDeleteFile отправляет видео через Telegram API с дополнительными параметрами и возвращает message_id
func SendVideoAndDeleteFile(videoConfig model.VideoConfig, messageID int64) (int64, error) {
	endpoint := fmt.Sprintf("%s/sendVideo", videoConfig.FullEndpoint)

	// Open the video file
	file, err := os.Open(videoConfig.FilePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Create a multipart form request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the video file to the form
	part, err := writer.CreateFormFile("video", filepath.Base(videoConfig.FilePath))
	if err != nil {
		return 0, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return 0, err
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
		return 0, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Parse the response
	var result struct {
		Ok     bool `json:"ok"`
		Result struct {
			MessageID int64 `json:"message_id"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	if !result.Ok {
		return 0, fmt.Errorf("failed to send video")
	}

	file.Close()

	// Добавляем задержку
	time.Sleep(100 * time.Millisecond)

	// Попытка удалить файл
	if err := os.Remove(videoConfig.FilePath); err != nil {
		fmt.Println("Failed to delete file:", err)
		return 0, err
	}

	err = DeleteMessage(videoConfig.ChatID, videoConfig.FullEndpoint, messageID)
	if err != nil {
		return 0, err
	}

	return result.Result.MessageID, nil
}

// SendMediaContent sends a group of media files (videos and photos) as a single message via Telegram API
func SendMediaContent(mediaConfig model.MediaContentConfig, messageID int64) ([]int64, error) {
	endpoint := fmt.Sprintf("%s/sendMediaGroup", mediaConfig.FullEndpoint)

	captionAdded := false

	// Prepare the payload
	var mediaItems []map[string]interface{}
	for _, video := range mediaConfig.VideosConfigs {
		mediaItem := map[string]interface{}{
			"type":               "video",
			"media":              fmt.Sprintf("attach://%s", filepath.Base(video.FilePath)),
			"supports_streaming": true,
		}
		if video.Duration != 0 {
			mediaItem["duration"] = strconv.Itoa(video.Duration)
		}
		if !captionAdded {
			caption := generateCaption(mediaConfig, video.Timing)
			mediaItem["caption"] = caption
			mediaItem["parse_mode"] = "HTML"
			captionAdded = true
		}
		mediaItems = append(mediaItems, mediaItem)
	}
	for _, photo := range mediaConfig.PhotosConfigs {
		mediaItem := map[string]interface{}{
			"type":  "photo",
			"media": fmt.Sprintf("attach://%s", filepath.Base(photo.FilePath)),
		}
		if !captionAdded {
			mediaItem["caption"] = generateCaption(mediaConfig, "")
			mediaItem["parse_mode"] = "HTML"
			captionAdded = true
		}
		mediaItems = append(mediaItems, mediaItem)
	}

	for _, audio := range mediaConfig.AudioConfigs {
		mediaItem := map[string]interface{}{
			"type":  "audio",
			"media": fmt.Sprintf("attach://%s", filepath.Base(audio.FilePath)),
		}
		if audio.Duration != 0 {
			mediaItem["duration"] = strconv.Itoa(audio.Duration)
		}
		if !captionAdded {
			mediaItem["caption"] = generateCaption(mediaConfig, "")
			mediaItem["parse_mode"] = "HTML"
			captionAdded = true
		}
		mediaItems = append(mediaItems, mediaItem)
	}

	// Serialize media items to JSON
	mediaItemsJSON, err := json.Marshal(mediaItems)
	if err != nil {
		return []int64{}, fmt.Errorf("failed to marshal media items: %v", err)
	}

	// Create a multipart request with each file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if err := writer.WriteField("chat_id", strconv.FormatInt(mediaConfig.ChatID, 10)); err != nil {
		return []int64{}, err
	}
	if mediaConfig.ReplyToMessageID != 0 {
		if err := writer.WriteField("reply_to_message_id", strconv.FormatInt(mediaConfig.ReplyToMessageID, 10)); err != nil {
			return []int64{}, err
		}
	}
	if err := writer.WriteField("media", string(mediaItemsJSON)); err != nil {
		return []int64{}, err
	}

	// Attach files to the request
	for _, video := range mediaConfig.VideosConfigs {
		file, err := os.Open(video.FilePath)
		if err != nil {
			return []int64{}, err
		}

		part, err := writer.CreateFormFile(filepath.Base(video.FilePath), filepath.Base(video.FilePath))
		if err != nil {
			file.Close()
			return []int64{}, err
		}
		if _, err := io.Copy(part, file); err != nil {
			file.Close()
			return []int64{}, err
		}
		file.Close()
	}

	for _, photo := range mediaConfig.PhotosConfigs {
		file, err := os.Open(photo.FilePath)
		if err != nil {
			return []int64{}, err
		}

		part, err := writer.CreateFormFile(filepath.Base(photo.FilePath), filepath.Base(photo.FilePath))
		if err != nil {
			file.Close()
			return []int64{}, err
		}
		if _, err := io.Copy(part, file); err != nil {
			file.Close()
			return []int64{}, err
		}
		file.Close()
	}

	for _, audio := range mediaConfig.AudioConfigs {
		file, err := os.Open(audio.FilePath)
		if err != nil {
			return []int64{}, err
		}

		part, err := writer.CreateFormFile(filepath.Base(audio.FilePath), filepath.Base(audio.FilePath))
		if err != nil {
			file.Close()
			return []int64{}, err
		}
		if _, err := io.Copy(part, file); err != nil {
			file.Close()
			return []int64{}, err
		}
		file.Close()
	}

	// Close the writer to finalize the request body
	if err := writer.Close(); err != nil {
		return []int64{}, err
	}

	// Send the request
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return []int64{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []int64{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []int64{}, err
	}

	// Log the response for debugging
	log.Println("Response Body:", string(bodyBytes))

	// Reset the response body so it can be read again
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Parse the response
	var result struct {
		Ok     bool `json:"ok"`
		Result []struct {
			MessageID int64 `json:"message_id"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []int64{}, err
	}

	if !result.Ok {
		return []int64{}, fmt.Errorf("failed to send media")
	}

	var messageIDs []int64

	for _, id := range result.Result {
		messageIDs = append(messageIDs, id.MessageID)
	}

	// Добавляем задержку
	time.Sleep(100 * time.Millisecond)

	// Попытка удалить файл

	// Clean up files after sending
	for _, video := range mediaConfig.VideosConfigs {
		if err := os.Remove(video.FilePath); err != nil {
			fmt.Println("Failed to delete file:", err)
			return []int64{}, err
		}
	}
	for _, photo := range mediaConfig.PhotosConfigs {
		if err := os.Remove(photo.FilePath); err != nil {
			fmt.Println("Failed to delete file:", err)
			return []int64{}, err
		}
	}
	for _, audio := range mediaConfig.AudioConfigs {
		if err := os.Remove(audio.FilePath); err != nil {
			fmt.Println("Failed to delete file:", err)
			return []int64{}, err
		}
	}

	// Delete the original message if needed
	err = DeleteMessage(mediaConfig.ChatID, mediaConfig.FullEndpoint, messageID)
	if err != nil {
		return []int64{}, err
	}

	return messageIDs, nil
}

// generateCaption generates a caption for a media item
func generateCaption(mediaConfig model.MediaContentConfig, timing string) string {
	caption := fmt.Sprintf("<a href='%s'>%s</a>", mediaConfig.Link, mediaConfig.Title)
	if timing != "" {
		caption += fmt.Sprintf("\n\n%s", timing)
	}
	if mediaConfig.Sender != "" {
		caption += fmt.Sprintf("\n\nSent by %s", mediaConfig.Sender)
	}
	return caption
}

// DeleteMessage удаляет сообщение по указанному chat_id и message_id
func DeleteMessage(chatID int64, fullEndpoint string, messageID int64) error {
	// Формируем URL для запроса к API
	endpoint := fmt.Sprintf("%s/deleteMessage", fullEndpoint)

	// Создаем запрос с параметрами chat_id и message_id
	data := url.Values{}
	data.Set("chat_id", strconv.FormatInt(chatID, 10))
	data.Set("message_id", strconv.FormatInt(messageID, 10))

	// Отправляем POST-запрос
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Выполняем запрос
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Обрабатываем ответ
	var result struct {
		Ok bool `json:"ok"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if !result.Ok {
		return fmt.Errorf("failed to delete message")
	}

	return nil
}
