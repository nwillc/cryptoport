package main

import (
	"fmt"
	"github.com/nwillc/cryptoport/pkg/externalapi/crypto"
	"github.com/nwillc/cryptoport/pkg/model"
)

func main() {
	p := model.Portfolio{
		Positions: []model.Position{
			{
				Currency: "BTC",
				Holding:  0.10345115,
			},
			{
				Currency: "ETH",
				Holding:  7.03096477,
			},
		},
	}

	appID := crypto.AppID("b0a17aba4805396e04667c975bcd5c3c0a6c480a")
	client, _ := crypto.NewClient(appID)
	service := crypto.NewNomicsCurrencyService(client)
	currencies := []crypto.Currency{"BTC", "ETH"}
	periods := []crypto.Period{"1d"}
	tickers, _ := service.Ticker(currencies, periods)

	values := p.Values(tickers)
	var total float64
	for k, v := range values {
		fmt.Printf("%20s %12.2f\n", k, v)
		total += v
	}
	fmt.Printf("%20s %12.2f\n", "Total:", total)
}
