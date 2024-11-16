package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gena97/telegram_bot/internal/app/config"
)

type currencyRates struct {
	Disclaimer string             `json:"disclaimer"`
	Date       string             `json:"date"`
	Timestamp  int64              `json:"timestamp"`
	Base       string             `json:"base"`
	Rates      map[string]float64 `json:"rates"`
}

type rankedPlayerStats struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			RankedGameModeStats struct {
				SquadFPP struct {
					CurrentTier struct {
						Tier    string `json:"tier"`
						SubTier string `json:"subTier"`
					} `json:"currentTier"`
					CurrentRankPoint int `json:"currentRankPoint"`
					BestTier         struct {
						Tier    string `json:"tier"`
						SubTier string `json:"subTier"`
					} `json:"bestTier"`
					BestRankPoint int     `json:"bestRankPoint"`
					KDA           float64 `json:"kda"`
					Top10Ratio    float64 `json:"top10Ratio"`
				} `json:"squad-fpp"`
			} `json:"rankedGameModeStats"`
		} `json:"attributes"`
	} `json:"data"`
}

func GetCurrencyRates() (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("https://www.cbr-xml-daily.ru/latest.js")
	if err != nil {
		return "", fmt.Errorf("ошибка при получении данных: %v", err)
	}
	defer resp.Body.Close()

	var data currencyRates
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("ошибка при декодировании JSON: %v", err)
	}

	// Получение курсов валют
	usdToRub := 1 / data.Rates["USD"]
	eurToRub := 1 / data.Rates["EUR"]
	rubToAmd := data.Rates["AMD"]
	date := data.Date

	// Формирование строки с курсами валют
	ratesString := fmt.Sprintf("Курсы ЦБ на %s\n🇺🇸USD:RUB   %.2f\n🇪🇺EUR:RUB   %.2f\n🇦🇲RUB:AMD  %.2f\n", date, usdToRub, eurToRub, rubToAmd)
	return ratesString, nil
}

func GetPubgStats() (string, error) {
	url := "https://api.pubg.com/shards/steam/players/" + config.ApiPubgAccID() + "/seasons/" + config.ApiPubgSeasonID() + "/ranked"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", config.ApiPubgKey())
	req.Header.Set("accept", "application/vnd.api+json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var stats rankedPlayerStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return "", err
	}

	// Извлечение нужной информации из JSON-ответа
	result := fmt.Sprintf("Статистика Миши в PUBG:\nТекущий тир: %s %s\nТекущие очки ранга: %d\nЛучший тир: %s %s\nЛучшие очки ранга: %d\nKDA: %.2f\nTop 10 Ratio: %.2f",
		stats.Data.Attributes.RankedGameModeStats.SquadFPP.CurrentTier.Tier,
		stats.Data.Attributes.RankedGameModeStats.SquadFPP.CurrentTier.SubTier,
		stats.Data.Attributes.RankedGameModeStats.SquadFPP.CurrentRankPoint,
		stats.Data.Attributes.RankedGameModeStats.SquadFPP.BestTier.Tier,
		stats.Data.Attributes.RankedGameModeStats.SquadFPP.BestTier.SubTier,
		stats.Data.Attributes.RankedGameModeStats.SquadFPP.BestRankPoint,
		stats.Data.Attributes.RankedGameModeStats.SquadFPP.KDA,
		stats.Data.Attributes.RankedGameModeStats.SquadFPP.Top10Ratio,
	)

	return result, nil
}
