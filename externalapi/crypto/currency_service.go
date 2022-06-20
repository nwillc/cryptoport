package crypto

import (
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/gslices"
	"github.com/nwillc/genfuncs/result"
)

const (
	tickerPath = "currencies/ticker"
)

// CurrencyService provides services to do with cryptocurrencies
type CurrencyService interface {
	Tickers(currencies container.GSlice[Currency], periods container.GSlice[Period]) *genfuncs.Result[container.GMap[Currency, *TickerInfo]]
}

// NomicsCurrencyService is a CurrencyService implemented using Nomics API.
type NomicsCurrencyService struct {
	client *Client
}

var _ CurrencyService = (*NomicsCurrencyService)(nil)

// NewNomicsCurrencyService creates a Nomics based CurrencyService.
func NewNomicsCurrencyService(client *Client) *NomicsCurrencyService {
	return &NomicsCurrencyService{
		client: client,
	}
}

// Tickers returns TickerInfo for the Currency list and Period list provides.
func (n NomicsCurrencyService) Tickers(currencies container.GSlice[Currency], periods container.GSlice[Period]) *genfuncs.Result[container.GMap[Currency, *TickerInfo]] {
	params := container.GMap[string, string]{
		"ids":      CurrencyList(currencies),
		"interval": PeriodList(periods),
	}
	tiResult := n.client.getTickerInfo(tickerPath, params)
	return result.Map(tiResult, func(list *container.GSlice[TickerInfo]) *genfuncs.Result[container.GMap[Currency, *TickerInfo]] {
		return genfuncs.NewResult(gslices.Associate(*list, func(t TickerInfo) (Currency, *TickerInfo) {
			return t.Currency, &t
		}))
	})
}
