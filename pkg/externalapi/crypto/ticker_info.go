package crypto

import (
	"strings"
)

type Currency string
type Price float64

type TickerInfo struct {
	Currency  Currency `json:"currency"`
	Price     Price    `json:"price,string"`
	Timestamp DateTime `json:"price_timestamp"`
}

func (c Currency) String() string {
	return string(c)
}

func CurrencyList(currencies []Currency) string {
	var strs []string
	for _, s := range currencies {
		strs = append(strs, string(s))
	}
	return strings.Join(strs, ",")
}

func PeriodList(periods []Period) string {
	var strs []string
	for _, s := range periods {
		strs = append(strs, string(s))
	}
	return strings.Join(strs, ",")
}
