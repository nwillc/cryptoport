package commands

import (
	"fmt"
	"github.com/fatih/color"
	crypto2 "github.com/nwillc/cryptoport/externalapi/crypto"
	model2 "github.com/nwillc/cryptoport/model"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"os"
)

var (
	colorGreen = color.New(color.FgGreen)
	colorRed   = color.New(color.FgRed)
	colorWhite = color.New(color.FgWhite)
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
	fileName := fmt.Sprintf("%s/%s", homeDir, model2.ConfFile)
	conf, err := model2.ReadConfig(fileName)
	if err != nil {
		panic(err)
	}
	client, err := crypto2.NewClient(conf.AppID)
	if err != nil {
		panic(err)
	}
	service := crypto2.NewNomicsCurrencyService(client)
	periods := []crypto2.Period{"1d"}
	var currencies []crypto2.Currency
	for _, position := range conf.Portfolio.Positions {
		currencies = append(currencies, position.Currency)
	}
	tickers, err := service.Tickers(currencies, periods)
	if err != nil {
		panic(err)
	}
	values := conf.Portfolio.Values(tickers)
	var total decimal.Decimal
	confValues := make(map[crypto2.Currency]decimal.Decimal)
	color := colorWhite
	for k, v := range values {
		if conf.Values != nil {
			hv, ok := (*conf.Values)[k.Currency]
			if ok {
				color = deltaColor(hv, v)
			}
		}
		_, _ = color.Printf("%20s %12s\n", k, v.StringFixed(2))
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
	_, _ = color.Printf("%20s %12s\n", "Total:", total.StringFixed(2))

	conf.Values = &confValues
	err = model2.WriteConfig(*conf, fileName)
	if err != nil {
		panic(err)
	}
}

func deltaColor(previous, current decimal.Decimal) *color.Color {
	switch {
	case previous.LessThan(current):
		return colorGreen
	case previous.GreaterThan(current):
		return colorRed
	default:
		return colorWhite
	}
}
