package model

import (
	"encoding/json"
	crypto2 "github.com/nwillc/cryptoport/externalapi/crypto"
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
	AppID     crypto2.AppID                         `json:"app_id"`
	Portfolio Portfolio                             `json:"portfolio"`
	Values    *map[crypto2.Currency]decimal.Decimal `json:"values,omitempty"`
}

// WriteConfig writes the Config given, as JSON to the filename given.
func WriteConfig(config Config, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	bytes, err := json.Marshal(&config)
	if err != nil {
		return err
	}
	_, err = file.Write(bytes)
	return err
}

// ReadConfig instantiates a Config from the JSON filename given.
func ReadConfig(filename string) (*Config, error) {
	readFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	conf := &Config{}
	err = json.Unmarshal(readFile, conf)
	return conf, err
}
