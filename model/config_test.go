package model

import (
	crypto2 "github.com/nwillc/cryptoport/externalapi/crypto"
	"github.com/nwillc/genfuncs"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteReadNoValues(t *testing.T) {
	conf := Config{
		AppID: "foo",
		Portfolio: Portfolio{
			Positions: []Position{
				{
					Currency: "BTC",
					Holding:  decimal.NewFromFloat(1.0),
				},
			},
		},
	}
	file := genfuncs.NewResultError(ioutil.TempFile("", "config"))
	require.True(t, file.Ok())
	fileName := file.OrEmpty().Name()
	t.Cleanup(func() {
		_ = os.Remove(fileName)
	})

	result := WriteConfig(conf, fileName)
	require.True(t, result.Ok())

	conf2 := ReadConfig(fileName)
	require.True(t, conf2.Ok())

	assert.Equal(t, conf.AppID, conf2.OrEmpty().AppID)
	assert.Equal(t, conf.Portfolio, conf2.OrEmpty().Portfolio)
	assert.Equal(t, 0, len(conf2.OrEmpty().Prices))
}

func TestWriteReadWithValues(t *testing.T) {
	price := decimal.NewFromFloat(40000.1)
	conf := Config{
		AppID: "foo",
		Portfolio: Portfolio{
			Positions: []Position{
				{
					Currency: "BTC",
					Holding:  decimal.NewFromFloat(1.0),
				},
			},
		},
		Prices: map[crypto2.Currency]decimal.Decimal{"BTC": price},
	}

	file := genfuncs.NewResultError(ioutil.TempFile("", "config"))
	require.True(t, file.Ok())
	fileName := file.OrEmpty().Name()
	t.Cleanup(func() {
		_ = os.Remove(fileName)
	})

	result := WriteConfig(conf, fileName)
	require.True(t, result.Ok())

	conf2 := ReadConfig(fileName)
	require.True(t, conf2.Ok())

	assert.Equal(t, conf.AppID, conf2.OrEmpty().AppID)
	assert.Equal(t, conf.Portfolio, conf2.OrEmpty().Portfolio)
	assert.Equal(t, price, conf2.OrEmpty().Prices["BTC"])
}
