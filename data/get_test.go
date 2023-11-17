package data

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCoinsRanking(t *testing.T) {
	servie := NewCoinMarketCap()
	res, err := servie.CoinsRanking(2, "market_cap")
	if err != nil {
		fmt.Println("Error ========>  " + err.Error())
	}
	a, err := json.MarshalIndent(*res, "", "    ")
	if err != nil {
		fmt.Println("Error ========>  " + err.Error())
	}
	// fmt.Printf("response =====> %+v", *res)
	fmt.Println(string(a))
}
