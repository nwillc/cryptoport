package crypto

import (
	"fmt"
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
	ticker, err := service.Tickers(currencies, periods)
	assert.NoError(t, err)
	fmt.Println(ticker)
}
