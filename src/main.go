package main

import (
	"trade_made_cron/src/client"
	"trade_made_cron/src/config"
	"trade_made_cron/src/controller"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.GetAPIConfig()
	client := client.NewTradeMade(cfg.ClientURI)
	ctrl := controller.NewController(client)

	res, err := ctrl.GetRequest()
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetHttpRequest"}).Error(err.Error())
		return
	}

	log.WithFields(log.Fields{"package": "httpServer", "method": "GetHttpRequest"}).Info(res)
}
