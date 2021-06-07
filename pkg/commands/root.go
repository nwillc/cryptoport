package commands

import (
	"fmt"
	"github.com/nwillc/cryptoport/pkg/externalapi/crypto"
	"github.com/nwillc/cryptoport/pkg/model"
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
	var total float64
	confValues := make(map[crypto.Currency]float64)
	color := colorWhite
	for k, v := range values {
		if conf.Values != nil {
			hv, ok := (*conf.Values)[k.Currency]
			if ok {
				color = deltaColor(hv, v)
			}
		}
		fmt.Printf("%s%20s %12.2f %s\n", color, k, v, colorReset)
		confValues[k.Currency] = v
		total += v
	}
	color = colorWhite
	if conf.Values != nil {
		var oldValue float64
		for _, v := range *conf.Values {
			oldValue += v
		}
		color = deltaColor(oldValue, total)
	}
	fmt.Printf("%s%20s %12.2f%s\n", color, "Total:", total, colorReset)

	conf.Values = &confValues
	err = model.WriteConfig(*conf, fileName)
	if err != nil {
		panic(err)
	}
}

func deltaColor(old, now float64) string {
	if old > now {
		return colorRed
	}
	if old < now {
		return colorGreen
	}
	return colorWhite
}
