package model

import (
	"github.com/nwillc/cryptoport/pkg/externalapi/crypto"
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
					Holding:  1,
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
					Holding:  1,
				},
			},
		},
		Values: &map[crypto.Currency]float64{
			"BTC": 40,
		},
	}

	err = WriteConfig(conf, file.Name())
	require.NoError(t, err)

	conf2, err := ReadConfig(file.Name())
	require.NoError(t, err)

	assert.Equal(t, conf.AppID, conf2.AppID)
	assert.Equal(t, conf.Portfolio, conf2.Portfolio)
	require.NotNil(t, conf2.Values)
	assert.Equal(t, *conf.Values, *conf2.Values)
}
