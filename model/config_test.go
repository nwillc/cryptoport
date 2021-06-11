package model

import (
	crypto2 "github.com/nwillc/cryptoport/externalapi/crypto"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteReadNoValues(t *testing.T) {
	file, err := ioutil.TempFile("", "config")
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = os.Remove(file.Name())
	})

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

	err = WriteConfig(conf, file.Name())
	require.NoError(t, err)

	conf2, err := ReadConfig(file.Name())
	require.NoError(t, err)

	assert.Equal(t, conf.AppID, conf2.AppID)
	assert.Equal(t, conf.Portfolio, conf2.Portfolio)
	assert.Nil(t, conf2.Values)
}

func TestWriteReadWithValues(t *testing.T) {
	file, err := ioutil.TempFile("", "config")
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = os.Remove(file.Name())
	})

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
		Values: &map[crypto2.Currency]decimal.Decimal{
			"BTC": decimal.NewFromFloat(40.0),
		},
	}

	err = WriteConfig(conf, file.Name())
	require.NoError(t, err)

	conf2, err := ReadConfig(file.Name())
	require.NoError(t, err)

	assert.Equal(t, conf.AppID, conf2.AppID)
	assert.Equal(t, conf.Portfolio, conf2.Portfolio)
	require.NotNil(t, conf2.Values)
	for k, v := range *conf.Values {
		assert.True(t, v.Equal((*conf2.Values)[k]))
	}
}
