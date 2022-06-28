package model

import (
	crypto2 "github.com/nwillc/cryptoport/externalapi/crypto"
	"github.com/nwillc/cryptoport/gjson"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/results"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"os"
)

const (
	// ConfFile name of the configuration file.
	ConfFile = ".crypto_port.json"
)

// Config persisted configuration.
type Config struct {
	AppID     crypto2.AppID                        `json:"app_id"`
	Portfolio Portfolio                            `json:"portfolio"`
	Prices    map[crypto2.Currency]decimal.Decimal `json:"prices,omitempty"`
}

// WriteConfig writes the Config given, as JSON to the filename given.
func WriteConfig(config Config, filename string) *genfuncs.Result[int] {
	result := gjson.Marshal(&config)
	return results.Map(result, func(bytes []byte) *genfuncs.Result[int] {
		file := genfuncs.NewResultError(os.Create(filename))
		return results.Map(file, func(f *os.File) *genfuncs.Result[int] {
			defer func() {
				_ = f.Close()
			}()
			return genfuncs.NewResultError(f.Write(bytes))
		})
	})
}

// ReadConfig instantiates a Config from the JSON filename given.
func ReadConfig(filename string) *genfuncs.Result[*Config] {
	readFile := genfuncs.NewResultError(ioutil.ReadFile(filename))
	return results.Map(readFile, func(bytes []byte) *genfuncs.Result[*Config] {
		return gjson.Unmarshal[Config](bytes)
	})
}
