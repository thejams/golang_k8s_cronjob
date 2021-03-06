//Package client provide the posibility to create clients to make http request to externals APIs
package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"trade_made_cron/src/entity"

	log "github.com/sirupsen/logrus"
)

type tradeMadeStruct struct {
	url string
}

//NewTradeMade initialice a new trade made controller
func NewTradeMade(url string) Client {
	log.SetFormatter(&log.JSONFormatter{})
	return &tradeMadeStruct{
		url: url,
	}
}

//HttpGetRequest makes an http get request to the public TradeMade API
func (c *tradeMadeStruct) HttpGetRequest() (interface{}, error) {
	res, err := http.Get(c.url)
	if err != nil {
		log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "HttpGetRequest"}).Error(err.Error())
		return nil, err
	}
	if res.StatusCode == 500 {
		log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "HttpGetRequest"}).Error("TradeMade client error 500")
		return nil, fmt.Errorf("TradeMade client error 500")
	}
	if res.StatusCode == 404 {
		log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "HttpGetRequest"}).Error("TradeMade client error 404")
		return nil, fmt.Errorf("TradeMade client error 404")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "HttpGetRequest"}).Error(err.Error())
		return nil, err
	}

	tradeMadeObj := entity.TradeMade{}
	err = json.Unmarshal(body, &tradeMadeObj)
	if err != nil {
		log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "Get"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "Get"}).Info("ok")
	return tradeMadeObj, nil
}
