package commands

import (
	"fmt"
	"github.com/fatih/color"
	crypto2 "github.com/nwillc/cryptoport/externalapi/crypto"
	"github.com/nwillc/cryptoport/model"
	"github.com/nwillc/genfuncs"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"os"
)

var (
	colorGreen = color.New(color.FgGreen)
	colorRed   = color.New(color.FgRed)
	colorWhite = color.New(color.FgWhite)
	oneHundred = decimal.NewFromFloat(100.0)
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
	homeDirResult := genfuncs.NewResultError(os.UserHomeDir())
	home := homeDirResult.MustGet()
	fileName := fmt.Sprintf("%s/%s", home, model.ConfFile)
	confResult := model.ReadConfig(fileName)
	conf := confResult.MustGet()
	clientResult := crypto2.NewClient(conf.AppID)
	client := clientResult.MustGet()
	service := crypto2.NewNomicsCurrencyService(client)
	periods := []crypto2.Period{"1d"}
	var currencies []crypto2.Currency
	for _, position := range conf.Portfolio.Positions {
		currencies = append(currencies, position.Currency)
	}
	tickersResult := service.Tickers(currencies, periods)
	tickers := tickersResult.MustGet()
	values := conf.Portfolio.Values(tickers)
	var total decimal.Decimal
	confValues := make(map[crypto2.Currency]decimal.Decimal)
	var color *color.Color
	var change decimal.Decimal
	for k, v := range values {
		if conf.Prices != nil {
			hv, ok := conf.Prices[k.Currency]
			if ok {
				color, change = delta(hv, v)
			}
		}
		_, _ = color.Printf("%20s %12s (%6s%%)\n", k, v.StringFixed(2), change.StringFixed(2))
		confValues[k.Currency] = v
		total = total.Add(v)
	}
	color = colorWhite
	if conf.Prices != nil {
		var oldValue decimal.Decimal
		for _, v := range conf.Prices {
			oldValue = oldValue.Add(v)
		}
		color, change = delta(oldValue, total)
	}
	_, _ = color.Printf("%20s %12s (%6s%%)\n", "Total:", total.StringFixed(2), change.StringFixed(2))

	conf.Prices = confValues
	result := model.WriteConfig(*conf, fileName)
	if !result.Ok() {
		panic(result.Error())
	}
}

func delta(previous, current decimal.Decimal) (*color.Color, decimal.Decimal) {
	change := percentChange(previous, current)
	switch {
	case change.GreaterThan(decimal.Zero):
		return colorGreen, change
	case change.LessThan(decimal.Zero):
		return colorRed, change
	default:
		return colorWhite, change
	}
}

func percentChange(previous, current decimal.Decimal) decimal.Decimal {
	delta := current.Sub(previous)
	return delta.Div(previous).Mul(oneHundred)
}
