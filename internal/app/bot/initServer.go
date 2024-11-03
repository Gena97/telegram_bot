package bot

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/Gena97/telegram_bot/internal/app/config"
)

func InitTelegramServer(cfg *config.Config) error {
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
		telegramCmd = exec.Command("../../utilities/telegram-server/telegram-bot-api.exe", "--api-id", cfg.TelegramServerID, "--api-hash", cfg.TelegramServerHash)

		// Запуск команды
		err = telegramCmd.Start()
		if err != nil {
			return fmt.Errorf("ошибка при запуске приложения: %v", err)
		}

		fmt.Println("Приложение telegram-bot-api.exe успешно запущено.")
	}
	return nil
}
