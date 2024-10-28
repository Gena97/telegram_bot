package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Gena97/telegram_bot/internal/config"
)

// Update - структура для получения обновлений
type Update struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int    `json:"message_id"`
		Text      string `json:"text"`
		Chat      struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

// Run запускает бота и обрабатывает обновления
func Run(cfg *config.Config) error {
	apiURL := cfg.BotUrl + cfg.TelegramToken

	// Начинаем цикл получения обновлений
	offset := 0
	for {
		updates, err := getUpdates(apiURL, offset)
		if err != nil {
			log.Printf("Error fetching updates: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		for _, update := range updates {
			offset = update.UpdateID + 1

			// Обрабатываем команду /start
			if update.Message.Text == "/start" {
				err := sendMessage(apiURL, update.Message.Chat.ID, "Welcome to the bot!")
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
			}
		}

		time.Sleep(1 * time.Second)
	}
}

// getUpdates получает обновления от Telegram API
func getUpdates(apiURL string, offset int) ([]Update, error) {
	endpoint := fmt.Sprintf("%s/getUpdates?offset=%d", apiURL, offset)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Ok     bool     `json:"ok"`
		Result []Update `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Result, nil
}

// sendMessage отправляет сообщение через Telegram API
func sendMessage(apiURL string, chatID int64, text string) error {
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
