package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetQuote() (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	baseURL := "http://api.forismatic.com/api/1.0/"
	url := fmt.Sprintf("%s?method=getQuote&format=json&lang=ru", baseURL)

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	quoteText, ok := data["quoteText"].(string)
	if !ok {
		return "", fmt.Errorf("цитата не найдена")
	}
	quoteAuthor, ok := data["quoteAuthor"].(string)
	if ok {
		quoteText += "\n\n" + quoteAuthor
	}

	return quoteText, nil
}
