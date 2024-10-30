package main

import (
	"github.com/Gena97/telegram_bot/internal/pkg/app"
)

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
