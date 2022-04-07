package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	AltaFinanceEarnURL = "https://alta.finance/api/earn"
)

type Earn struct {
	Tvl    float64 `json:"tvl"`
	MaxAPR float64 `json:"maxAPR"`
	Data   []struct {
		ContractAddress string  `json:"contractAddress"`
		AprBase         float64 `json:"aprBase"`
		AprBonus        float64 `json:"aprBonus"`
		SymbolBase      string  `json:"symbolBase"`
		SymbolBonus     string  `json:"symbolBonus"`
		Tvl             float64 `json:"tvl"`
		Network         string  `json:"network"`
		Link            string  `json:"link"`
	} `json:"data"`
}

func GetEarn() (result Earn, err error) {

	req, err := http.NewRequest("GET", AltaFinanceEarnURL, nil)
	if err != nil {
		return result, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(results, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
