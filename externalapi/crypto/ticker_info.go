package crypto

import (
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/sequences"
	"github.com/shopspring/decimal"
)

// Currency name, BTC, ETH etc.
type Currency string

// TickerInfo contains a Price for a Currency at a given Timestamp.
type TickerInfo struct {
	Currency  Currency        `json:"currency"`
	Price     decimal.Decimal `json:"price"`
	Timestamp DateTime        `json:"price_timestamp"`
}

// String implements fmt.Stringer for Currency.
func (c Currency) String() string {
	return string(c)
}

// CurrencyList formats a list of Currency into a string.
func CurrencyList(currencies container.GSlice[Currency]) string {
	return sequences.JoinToString[Currency](currencies, genfuncs.StringerToString[Currency](), ",", "", "")
}

// PeriodList formats a list of Period into a string.
func PeriodList(periods container.GSlice[Period]) string {
	return sequences.JoinToString[Period](periods, func(p Period) string { return string(p) }, ",", "", "")
}
