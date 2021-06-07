package crypto

import (
	"strings"
)

// Currency name, BTC, ETH etc.
type Currency string

// Price representation.
type Price float64

// TickerInfo contains a Price for a Currency at a given Timestamp.
type TickerInfo struct {
	Currency  Currency `json:"currency"`
	Price     Price    `json:"price,string"`
	Timestamp DateTime `json:"price_timestamp"`
}

// String implements fmt.Stringer for Currency.
func (c Currency) String() string {
	return string(c)
}

// CurrencyList formats a list of Currency into a string.
func CurrencyList(currencies []Currency) string {
	var strs []string
	for _, s := range currencies {
		strs = append(strs, string(s))
	}
	return strings.Join(strs, ",")
}

// PeriodList formats a list of Period into a string.
func PeriodList(periods []Period) string {
	var strs []string
	for _, s := range periods {
		strs = append(strs, string(s))
	}
	return strings.Join(strs, ",")
}
