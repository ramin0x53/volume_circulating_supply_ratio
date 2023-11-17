package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type CoinMarketCap struct {
	Addr string
}

func NewCoinMarketCap() *CoinMarketCap {
	return &CoinMarketCap{Addr: "https://api.coinmarketcap.com"}
}

func (c *CoinMarketCap) CoinsRanking(limit int64, sortBy string) (*CoinRankingResponse, error) {
	url := fmt.Sprintf(c.Addr+"/data-api/v3/cryptocurrency/listing?start=1&limit=%d&sortBy=%s&sortType=desc&convert=USD,BTC,ETH&cryptoType=all&tagType=all&audited=false&aux=ath,atl,high24h,low24h,num_market_pairs,cmc_rank,date_added,max_supply,circulating_supply,total_supply,volume_7d,volume_30d,self_reported_circulating_supply,self_reported_market_cap", limit, sortBy)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Host", "api.coinmarketcap.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/118.0")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("status code: " + res.Status)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result CoinRankingResponse
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
