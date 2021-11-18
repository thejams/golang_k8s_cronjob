package controller

import (
	"fmt"
	"trade_made_cron/src/client"
	"trade_made_cron/src/entity"

	log "github.com/sirupsen/logrus"
)

//Controller main interface for comunicate with the client interface
type Controller interface {
	GetRequest() (interface{}, error)
}

type controller struct {
	client client.Client
}

//NewController create a new controller
func NewController(client client.Client) Controller {
	log.SetFormatter(&log.JSONFormatter{})
	return &controller{
		client: client,
	}
}

//GetRequest makes an http get request via client
func (s *controller) GetRequest() (interface{}, error) {
	res, err := s.client.HttpGetRequest()
	if err != nil {
		log.WithFields(log.Fields{"package": "controller", "method": "GetRequest"}).Error(err.Error())
		return nil, err
	}

	tradeMadeObj, ok := res.(entity.TradeMade)
	if !ok {
		log.WithFields(log.Fields{"package": "controller", "method": "GetRequest"}).Error("TradeMade converstion fail")
		return nil, fmt.Errorf("TradeMade converstion fail")
	}
	if tradeMadeObj.Endpoint == "" || len(tradeMadeObj.Quotes) == 0 {
		log.WithFields(log.Fields{"package": "controller", "method": "GetRequest"}).Error("no TradeMade response received")
		return nil, fmt.Errorf("no TradeMade response received")
	}

	log.WithFields(log.Fields{"package": "controller", "method": "GetRequest"}).Info("ok")
	return tradeMadeObj, nil
}
