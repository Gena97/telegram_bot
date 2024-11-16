package model

import "github.com/tidwall/gjson"

const MemniyRayChatID = -1001290328010
const MemniyRayTag = "memniy_ray"

type TelegramBot struct {
	TelegramBotChan chan gjson.Result
	FullEndpoint    string
	Token           string
}

const (
	ChatTypePrivate = "private"
)

type Message struct {
	RawMsg           gjson.Result
	Text             string
	ChatID           int64
	ChatType         string
	ReplyToMessageID int64
	FromFirstName    string
	MessageID        int64
	FromID           int64
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
	Thumbnail        []byte
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

type User struct {
	UserID    int64
	FirstName string
	LastName  string
	Username  string
	IsAdmin   bool
}
