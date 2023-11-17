package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"volume_circulating_supply_ratio/data"

	"github.com/gin-gonic/gin"
)

type Coin struct {
	Name                         string  `json:"name"`
	Symbol                       string  `json:"symbol"`
	QuoteName                    string  `json:"quoteName"`
	MarketCap                    float64 `json:"marketCap"`
	CirculatingSupply            float64 `json:"circulatingSupply"`
	Volume                       float64 `json:"volume"`
	VolumeCirculatingSupplyRatio float64 `json:"volumeCirculatingSupplyRatio"`
}

func Calculate(info *data.CoinRankingResponse, volumeDuration string, quoteName string) (*[]Coin, error) {
	coins := []Coin{}
	for _, coin := range (*info).Data.CryptoCurrencyList {
		for _, quote := range coin.Quotes {
			if quote.Name == quoteName {
				newCoin := Coin{}
				newCoin.Symbol = coin.Symbol
				newCoin.QuoteName = quote.Name
				newCoin.Name = coin.Name
				newCoin.MarketCap = quote.MarketCap
				switch volumeDuration {
				case "1d":
					newCoin.Volume = quote.Volume24h
					newCoin.CirculatingSupply = coin.CirculatingSupply * quote.Price
					newCoin.VolumeCirculatingSupplyRatio = newCoin.Volume / (newCoin.CirculatingSupply)
				case "7d":
					newCoin.Volume = quote.Volume7d
					newCoin.CirculatingSupply = coin.CirculatingSupply * quote.Price
					newCoin.VolumeCirculatingSupplyRatio = newCoin.Volume / (newCoin.CirculatingSupply)
				case "30d":
					newCoin.Volume = quote.Volume30d
					newCoin.CirculatingSupply = coin.CirculatingSupply * quote.Price
					newCoin.VolumeCirculatingSupplyRatio = newCoin.Volume / (newCoin.CirculatingSupply)
				default:
					return nil, errors.New("wrong volume duration")
				}
				coins = append(coins, newCoin)
			}
		}
	}
	return &coins, nil
}

func Handler(c *gin.Context) {
	limit := c.DefaultQuery("limit", "100")
	//can be 1d, 7d, 30d
	duration := c.DefaultQuery("duration", "1d")
	quoteName := strings.ToUpper(c.DefaultQuery("quoteName", "USD"))
	sortBy := c.DefaultQuery("sortBy", "market_cap")
	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := data.NewCoinMarketCap()
	response, err := service.CoinsRanking(limitInt, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	coins, err := Calculate(response, duration, quoteName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quicksortCoins(*coins))
	// c.JSON(http.StatusOK, coins)
}

func quicksortCoins(coins []Coin) []Coin {
	if len(coins) <= 1 {
		return coins
	}

	pivotIndex := len(coins) / 2
	pivot := coins[pivotIndex]

	var less, equal, greater []Coin
	for _, coin := range coins {
		switch {
		case coin.VolumeCirculatingSupplyRatio < pivot.VolumeCirculatingSupplyRatio:
			less = append(less, coin)
		case coin.VolumeCirculatingSupplyRatio == pivot.VolumeCirculatingSupplyRatio:
			equal = append(equal, coin)
		case coin.VolumeCirculatingSupplyRatio > pivot.VolumeCirculatingSupplyRatio:
			greater = append(greater, coin)
		}
	}

	sortedCoins := append(append(quicksortCoins(less), equal...), quicksortCoins(greater)...)
	return sortedCoins
}

func main() {
	server := gin.Default()
	server.GET("/calculate", Handler)
	server.Run()
}
