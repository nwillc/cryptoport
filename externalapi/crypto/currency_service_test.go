package crypto

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewNomicsCurrencyService(t *testing.T) {
	var cs CurrencyService = NewNomicsCurrencyService(nil)
	require.NotNil(t, cs)
}

func Test_nomics_Ticker(t *testing.T) {
	client := MightSkipIntegrationTest(t)
	service := NewNomicsCurrencyService(client)
	require.NotNil(t, service)
	currencies := []Currency{"BTC", "ETH"}
	periods := []Period{"1d"}
	tickers := service.Tickers(currencies, periods)
	assert.True(t, tickers.Ok())
	tickers.OrEmpty().ForEach(func(c Currency, ti *TickerInfo) {
		t.Log(c, *ti)
	})
}
