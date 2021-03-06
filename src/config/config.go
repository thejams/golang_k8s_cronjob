package config

import (
	"os"
	"strings"
	"trade_made_cron/src/entity"
)

//GetAPIConfig returns config struct for the app
func GetAPIConfig() *entity.APPConfig {
	api_key := os.Getenv("API_KEY")
	currencies := os.Getenv("CURRENCIES")
	if len(strings.TrimSpace(currencies)) == 0 {
		currencies = "EURUSD,GBPUSD"
	}
	if len(strings.TrimSpace(api_key)) == 0 {
		api_key = "12345"
	}
	url := "https://marketdata.tradermade.com/api/v1/live?currency=" + currencies + "&api_key=" + api_key

	return &entity.APPConfig{
		ClientURI: url,
	}
}
