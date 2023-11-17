package data

import "time"

type Quote struct {
	Name                     string  `json:"name"`
	Price                    float64 `json:"price"`
	Volume24h                float64 `json:"volume24h"`
	Volume7d                 float64 `json:"volume7d"`
	Volume30d                float64 `json:"volume30d"`
	MarketCap                float64 `json:"marketCap"`
	SelfReportedMarketCap    float64 `json:"selfReportedMarketCap"`
	PercentChange1h          float64 `json:"percentChange1h"`
	PercentChange24h         float64 `json:"percentChange24h"`
	PercentChange7d          float64 `json:"percentChange7d"`
	LastUpdated              string  `json:"lastUpdated"`
	PercentChange30d         float64 `json:"percentChange30d"`
	PercentChange60d         float64 `json:"percentChange60d"`
	PercentChange90d         float64 `json:"percentChange90d"`
	FullyDilluttedMarketCap  float64 `json:"fullyDilluttedMarketCap"`
	MarketCapByTotalSupply   float64 `json:"marketCapByTotalSupply"`
	Dominance                float64 `json:"dominance"`
	Turnover                 float64 `json:"turnover"`
	YtdPriceChangePercentage float64 `json:"ytdPriceChangePercentage"`
	PercentChange1y          float64 `json:"percentChange1y"`
}

type CryptoCurrency struct {
	ID                            int64     `json:"id"`
	Name                          string    `json:"name"`
	Symbol                        string    `json:"symbol"`
	Slug                          string    `json:"slug"`
	CmcRank                       int64     `json:"cmcRank"`
	MarketPairCount               int64     `json:"marketPairCount"`
	CirculatingSupply             float64   `json:"circulatingSupply"`
	SelfReportedCirculatingSupply float64   `json:"selfReportedCirculatingSupply"`
	TotalSupply                   float64   `json:"totalSupply"`
	MaxSupply                     float64   `json:"maxSupply"`
	Ath                           float64   `json:"ath"`
	Atl                           float64   `json:"atl"`
	High24h                       float64   `json:"high24h"`
	Low24h                        float64   `json:"low24h"`
	IsActive                      int64     `json:"isActive"`
	LastUpdated                   time.Time `json:"lastUpdated"`
	DateAdded                     time.Time `json:"dateAdded"`
	Quotes                        []Quote   `json:"quotes"`
	IsAudited                     bool      `json:"isAudited"`
	// AuditInfoList                 []string  `json:"auditInfoList"`
	Badges []int64 `json:"badges"`
}

type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    string    `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      string    `json:"elapsed"`
	CreditCount  int64     `json:"credit_count"`
}

type Data struct {
	CryptoCurrencyList []CryptoCurrency `json:"cryptoCurrencyList"`
	TotalCount         string           `json:"totalCount"`
}

type CoinRankingResponse struct {
	Data   Data   `json:"data"`
	Status Status `json:"status"`
}
