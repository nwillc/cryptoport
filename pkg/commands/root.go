package commands

import (
	"fmt"
	"github.com/nwillc/cryptoport/pkg/externalapi/crypto"
	"github.com/nwillc/cryptoport/pkg/model"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"os"
)

const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorWhite = "\033[37m"
	colorReset = "\033[0m"
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "cryptoport",
	Short: "A simple crypto portfolio status cli",
	Long:  "A simple crypto portfolio status cli that reports the value of your portfolio",
	Args:  cobra.ExactArgs(0),
	Run:   portfolio,
}

func portfolio(_ *cobra.Command, _ []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("%s/%s", homeDir, model.ConfFile)
	conf, err := model.ReadConfig(fileName)
	if err != nil {
		panic(err)
	}
	client, err := crypto.NewClient(conf.AppID)
	if err != nil {
		panic(err)
	}
	service := crypto.NewNomicsCurrencyService(client)
	periods := []crypto.Period{"1d"}
	var currencies []crypto.Currency
	for _, position := range conf.Portfolio.Positions {
		currencies = append(currencies, position.Currency)
	}
	tickers, err := service.Tickers(currencies, periods)
	if err != nil {
		panic(err)
	}
	values := conf.Portfolio.Values(tickers)
	var total decimal.Decimal
	confValues := make(map[crypto.Currency]decimal.Decimal)
	color := colorWhite
	for k, v := range values {
		if conf.Values != nil {
			hv, ok := (*conf.Values)[k.Currency]
			if ok {
				color = deltaColor(hv, v)
			}
		}
		fmt.Printf("%s%20s %12s%s\n", color, k, v.StringFixed(2), colorReset)
		confValues[k.Currency] = v
		total = total.Add(v)
	}
	color = colorWhite
	if conf.Values != nil {
		var oldValue decimal.Decimal
		for _, v := range *conf.Values {
			oldValue = oldValue.Add(v)
		}
		color = deltaColor(oldValue, total)
	}
	fmt.Printf("%s%20s %12s%s\n", color, "Total:", total.StringFixed(2), colorReset)

	conf.Values = &confValues
	err = model.WriteConfig(*conf, fileName)
	if err != nil {
		panic(err)
	}
}

func deltaColor(old, now decimal.Decimal) string {
	switch {
	case old.LessThan(now):
		return colorGreen
	case old.GreaterThan(now):
		return colorRed
	default:
		return colorWhite
	}
}
