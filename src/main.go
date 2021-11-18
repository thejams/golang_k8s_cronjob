package main

import (
	"os"
	"strings"
	"trade_made_cron/src/client"
	"trade_made_cron/src/controller"

	log "github.com/sirupsen/logrus"
)

func main() {
	api_key := os.Getenv("API_KEY")
	currencies := os.Getenv("CURRENCIES")
	if len(strings.TrimSpace(currencies)) == 0 {
		currencies = "EURUSD,GBPUSD"
	}
	if len(strings.TrimSpace(api_key)) == 0 {
		api_key = "12345"
	}
	url := "https://marketdata.tradermade.com/api/v1/live?currency=" + currencies + "&api_key=" + api_key

	client := client.NewTradeMade(url)
	ctrl := controller.NewController(client)

	res, err := ctrl.GetRequest()
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetHttpRequest"}).Error(err.Error())
		return
	}
	log.WithFields(log.Fields{"package": "httpServer", "method": "GetHttpRequest"}).Info(res)
}
