package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/tidwall/gjson"
)

// RunBot запускает бота и возвращает канал с обновлениями
func GetTelegramBot(telegramBotEndpoint, telegramBotToken string) (model.TelegramBot, error) {
	apiURL := fmt.Sprintf("%s%s", telegramBotEndpoint, telegramBotToken)
	updatesChan := make(chan gjson.Result)

	go func() {
		var offset int64
		defer close(updatesChan)
		for {
			updates, err := GetUpdates(apiURL, offset, []string{"message", "message_reaction"})
			if err != nil {
				log.Printf("Error fetching updates: %v", err)
				time.Sleep(3 * time.Second)
				continue
			}

			for _, update := range updates {
				log.Println(update)
				offset = update.Get("update_id").Int() + 1
				updatesChan <- update // отправка обновления в канал
			}

			time.Sleep(1 * time.Second)
		}
	}()

	tgBot := model.TelegramBot{
		TelegramBotChan: updatesChan,
		FullEndpoint:    apiURL,
	}

	return tgBot, nil
}
