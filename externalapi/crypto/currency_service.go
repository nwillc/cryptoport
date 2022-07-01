package crypto

import (
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/maps"
	"github.com/nwillc/genfuncs/container/sequences"
	"github.com/nwillc/genfuncs/results"
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
	return results.Map(
		n.client.getTickerInfo(tickerPath, params),
		func(list *container.GSlice[TickerInfo]) *genfuncs.Result[container.GMap[Currency, *TickerInfo]] {
			return sequences.Associate[TickerInfo, Currency, *TickerInfo](
				*list,
				func(t TickerInfo) *genfuncs.Result[*maps.Entry[Currency, *TickerInfo]] {
					return genfuncs.NewResult(maps.NewEntry(t.Currency, &t))
				})
		})
}
