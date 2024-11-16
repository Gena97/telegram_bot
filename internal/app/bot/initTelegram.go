package bot

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/Gena97/telegram_bot/internal/app/config"
	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/tidwall/gjson"
)

func InitTelegramServer() error {
	cmd := exec.Command("tasklist")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {

		return fmt.Errorf("ошибка выполнения команды tasklist: %v", err)
	}

	// Проверка, содержится ли имя процесса в выводе команды
	output := out.String()
	var telegramCmd *exec.Cmd
	if strings.Contains(output, "telegram-bot-api.exe") {
		fmt.Println("Приложение telegram-bot-api.exe запущено.")
	} else {
		fmt.Println("Приложение telegram-bot-api.exe не запущено. Запускаю...")

		// Команда для запуска приложения с параметрами
		telegramCmd = exec.Command("../../utilities/telegram-server/telegram-bot-api.exe", "--api-id", config.TelegramServerID(), "--api-hash", config.TelegramServerHash(), "--http-port", "8888")

		// Запуск команды
		err = telegramCmd.Start()
		if err != nil {
			return fmt.Errorf("ошибка при запуске приложения: %v", err)
		}

		fmt.Println("Приложение telegram-bot-api.exe успешно запущено.")
	}
	return nil
}

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
		Token:           telegramBotToken,
	}

	return tgBot, nil
}
