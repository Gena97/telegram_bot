package service

import "strings"

// DetectContentType checks the content type of the provided URL and returns it as a string.
func DetectContentType(url string) string {
	switch {
	case strings.Contains(url, "youtu"):
		return "youtube"
	case strings.Contains(url, "https://www.x.com"), strings.Contains(url, "https://x.com"), strings.Contains(url, "twitter.com"):
		return "twitter"
	case strings.Contains(url, "inst"):
		return "instagram"
	default:
		return ""
	}
}

// isCommand проверяет, является ли текст команды.
func IsCommand(text string) bool {
	return strings.HasPrefix(text, "/")
}

func ParseCommandArgs(text string) []string {
	return strings.Split(text, " ")
}
