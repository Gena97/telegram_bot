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

type PhotoConfig struct {
	FilePath         string
	Title            string
	FullEndpoint     string
	ChatID           int64
	ReplyToMessageID int64
	PhotoURL         string
	Sender           string
}

type AudioConfig struct {
	FilePath         string
	Title            string
	FullEndpoint     string
	ChatID           int64
	ReplyToMessageID int64
	AudioURL         string
	Sender           string
	Duration         int
}

type MediaContentConfig struct {
	VideosConfigs    []VideoConfig
	PhotosConfigs    []PhotoConfig
	AudioConfigs     []AudioConfig
	Title            string
	Link             string
	FullEndpoint     string
	ChatID           int64
	ReplyToMessageID int64
	Sender           string
}
