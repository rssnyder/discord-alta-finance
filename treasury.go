package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	AltaFinanceTreasuryURL = "https://alta.finance/api/treasury"
)

type Treasury struct {
	Address []string `json:"address"`
	Wallet  struct {
		Response struct {
			Balance float64 `json:"balance"`
			Data    []struct {
				Asset       string      `json:"asset"`
				Value       interface{} `json:"value"`
				UsdValue    interface{} `json:"usdValue"`
				UsdPrice    float64     `json:"usdPrice"`
				Network     string      `json:"network"`
				RawContract struct {
					Address interface{} `json:"address"`
					Decimal int         `json:"decimal"`
				} `json:"rawContract"`
				Metadata struct {
					Decimals int    `json:"decimals"`
					Logo     string `json:"logo"`
					Name     string `json:"name"`
					Symbol   string `json:"symbol"`
				} `json:"metadata,omitempty"`
			} `json:"data"`
		} `json:"response"`
	} `json:"wallet"`
	Vault struct {
		Response struct {
			Balance float64 `json:"balance"`
			Data    []struct {
				Asset       string      `json:"asset"`
				Value       interface{} `json:"value"`
				UsdValue    interface{} `json:"usdValue"`
				UsdPrice    float64     `json:"usdPrice"`
				Network     string      `json:"network"`
				RawContract struct {
					Address interface{} `json:"address"`
					Decimal int         `json:"decimal"`
				} `json:"rawContract"`
				Metadata struct {
					Logo string `json:"logo"`
				} `json:"metadata"`
			} `json:"data"`
		} `json:"response"`
	} `json:"vault"`
	Exchanges struct {
		Success    bool `json:"success"`
		StatusCode int  `json:"statusCode"`
		Response   struct {
			Balance         int           `json:"balance"`
			BalanceEthereum int           `json:"balanceEthereum"`
			BalancePolygon  int           `json:"balancePolygon"`
			Data            []interface{} `json:"data"`
		} `json:"response"`
	} `json:"exchanges"`
	Funds struct {
		Success    bool `json:"success"`
		StatusCode int  `json:"statusCode"`
		Response   struct {
			Balance         int `json:"balance"`
			BalancePolygon  int `json:"balancePolygon"`
			BalanceEthereum int `json:"balanceEthereum"`
			Data            []struct {
				Balance int    `json:"balance"`
				Title   string `json:"title"`
				Data    []struct {
					Asset       string      `json:"asset"`
					Value       interface{} `json:"value"`
					UsdValue    interface{} `json:"usdValue"`
					UsdPrice    int         `json:"usdPrice"`
					Network     string      `json:"network"`
					RawContract struct {
						Address string `json:"address"`
						Decimal int    `json:"decimal"`
					} `json:"rawContract"`
					Metadata struct {
						Decimals int    `json:"decimals"`
						Name     string `json:"name"`
						Symbol   string `json:"symbol"`
						Logo     string `json:"logo"`
					} `json:"metadata"`
				} `json:"data"`
			} `json:"data"`
		} `json:"response"`
	} `json:"funds"`
	Earn struct {
		Success    bool `json:"success"`
		StatusCode int  `json:"statusCode"`
		Response   struct {
			Balance float64 `json:"balance"`
			Data    []struct {
				Asset       string      `json:"asset"`
				Value       interface{} `json:"value"`
				UsdValue    interface{} `json:"usdValue"`
				UsdPrice    int         `json:"usdPrice"`
				Network     string      `json:"network"`
				RawContract struct {
					Address string `json:"address"`
					Decimal int    `json:"decimal"`
				} `json:"rawContract"`
				Metadata struct {
					Decimals int    `json:"decimals"`
					Name     string `json:"name"`
					Symbol   string `json:"symbol"`
					Logo     string `json:"logo"`
				} `json:"metadata"`
			} `json:"data"`
		} `json:"response"`
	} `json:"earn"`
	Balance  float64 `json:"balance"`
	Networks struct {
		Balance  float64 `json:"balance"`
		Ethereum float64 `json:"ethereum"`
		Polygon  float64 `json:"polygon"`
	} `json:"networks"`
}

func GetTreasury() (result Treasury, err error) {

	req, err := http.NewRequest("GET", AltaFinanceTreasuryURL, nil)
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
