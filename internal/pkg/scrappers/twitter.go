package scrappers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/Gena97/telegram_bot/internal/service"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

func GetTwitterScrapper(login, password string) (*twitterscraper.Scraper, error) {
	// Twitter scraper initialization and login
	twitterScrapper := twitterscraper.New()

	// Attempt to load cookies from file
	loadedCookies, err := loadCookiesFromFile("../../cookiesTwitter.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("Cookies file not found, proceeding with login.")
		} else {
			log.Printf("Error loading cookies: %s", err)
		}
	} else {
		// Set loaded cookies to the scraper
		twitterScrapper.SetCookies(loadedCookies)
		log.Printf("Loaded cookies from file.")
	}

	// Check if already logged in
	if twitterScrapper.IsLoggedIn() {
		log.Printf("Twitter scraper logged in successfully!")
		return twitterScrapper, nil
	}

	// Attempt login if not logged in
	log.Printf("Attempting to log in to Twitter...")
	if err := twitterScrapper.Login(login, password); err != nil {
		return nil, fmt.Errorf("error logging in to Twitter scraper: %w", err)
	}
	log.Printf("Twitter scraper logged in successfully!")

	// Save updated cookies after successful login
	updatedCookies := twitterScrapper.GetCookies()
	if saveErr := saveCookiesToFile(updatedCookies, "cookiesTwitter.json"); saveErr != nil {
		log.Printf("Error saving cookies: %s", saveErr)
	} else {
		log.Printf("Cookies saved successfully to %s", "cookiesTwitter.json")
	}

	return twitterScrapper, nil
}

func loadCookiesFromFile(filePath string) ([]*http.Cookie, error) {
	var cookies []*http.Cookie

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, os.ErrNotExist
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(file, &cookies); err != nil {
		return nil, err
	}

	return cookies, nil
}

func saveCookiesToFile(cookies []*http.Cookie, filePath string) error {
	js, err := json.Marshal(cookies)
	if err != nil {
		return err
	}

	if err = os.WriteFile(filePath, js, fs.FileMode(0644)); err != nil {
		return err
	}
	return nil
}
func DownloadContentTwitter(tweetURL string, scrapper *twitterscraper.Scraper) (model.MediaContentConfig, error) {
	twitNumber := extractNumber(tweetURL)
	var twitterContent model.MediaContentConfig

	// Retrieve tweet details
	tweet, err := scrapper.GetTweet(twitNumber)
	if err != nil {
		log.Printf("Error retrieving tweet video")
		return model.MediaContentConfig{}, err
	}

	// Check for media content in the tweet
	if len(tweet.Videos) == 0 && len(tweet.Photos) == 0 {
		return model.MediaContentConfig{}, nil
	}

	// Define a directory to store downloaded media
	directory := "../../downloads" // Adjust as needed

	// Download videos
	for _, video := range tweet.Videos {
		filePath, err := service.DownloadMedia(video.URL, directory)
		if err != nil {
			log.Printf("Error downloading video: %v", err)
			continue
		}

		duration, err := service.GetDuration(filePath)
		if err != nil {
			log.Printf("Ошибка при получении продолжительности файла: %v", err)
			return model.MediaContentConfig{}, err
		}

		videoConfig := model.VideoConfig{FilePath: filePath, Title: tweet.Text, VideoURL: video.URL, Duration: duration}

		twitterContent.VideosConfigs = append(twitterContent.VideosConfigs, videoConfig)
	}

	// Download photos
	for _, photo := range tweet.Photos {
		filePath, err := service.DownloadMedia(photo.URL, directory)
		if err != nil {
			log.Printf("Error downloading photo: %v", err)
			continue
		}
		photoConfig := model.PhotoConfig{FilePath: filePath, Title: tweet.Text, PhotoURL: photo.URL}

		twitterContent.PhotosConfigs = append(twitterContent.PhotosConfigs, photoConfig)
	}

	twitterContent.Link = tweetURL
	twitterContent.Title = tweet.Text

	return twitterContent, nil
}

func extractNumber(url string) string {
	parts := strings.Split(url, "/") // Разбиваем ссылку по символу '/'
	lastPart := parts[len(parts)-1]  // Берем последнюю часть

	// Разбиваем последнюю часть по символу '?'
	numberParts := strings.Split(lastPart, "?")
	number := numberParts[0] // Берем номер до символа '?'

	return number
}
