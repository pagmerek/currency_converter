package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type exchange_rate struct {
	Table    string `json:"table"`
	Currency string `json:"currency"`
	Code     string `json:"code"`
	Rates    []struct {
		No            string  `json:"no"`
		EffectiveDate string  `json:"effectiveDate"`
		Mid           float32 `json:"mid"`
	} `json:"rates"`
}

func getExchangeRate(currency_id string) float32 {
	url := "http://api.nbp.pl/api/exchangerates/rates/a/" + currency_id + "/?format=json"
	convertClient := http.Client{
		Timeout: time.Second * 2,
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("User-Agent", "currency_converter")

	if err != nil {
		log.Fatal(err)
	}
	response, err := convertClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	result := exchange_rate{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Rates[0].Mid
}
