package model

import "github.com/tidwall/gjson"

type TelegramBot struct {
	TelegramBotChan chan gjson.Result
	FullEndpoint    string
}

type VideoConfig struct {
	FilePath         string
	Title            string
	Duration         int
	Timing           string
	FullEndpoint     string
	ChatID           int64
	ReplyToMessageID int64
	VideoURL         string
	Sender           string
}
