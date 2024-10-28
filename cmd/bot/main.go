package main

import (
	"log"

	"github.com/Gena97/telegram_bot/internal/bot"
	"github.com/Gena97/telegram_bot/internal/config"
)

func main() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := bot.Run(cfg); err != nil {
		log.Fatalf("error running bot: %v", err)
	}
}
