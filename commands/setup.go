package commands

import (
	"fmt"
	"github.com/nwillc/cryptoport/externalapi/crypto"
	"github.com/nwillc/cryptoport/model"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(setupCmd)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup your portfolio configuration.",
	Long:  "Setup your portfolio configuration and save to " + model.ConfFile,
	Run:   setup,
}

func setup(_ *cobra.Command, _ []string) {
	conf := model.Config{
		AppID:     "",
		Portfolio: model.Portfolio{},
	}
	for {
		conf.AppID = crypto.AppID(readString("Enter Nomics API_ID: "))
		if conf.AppID == "" {
			_, _ = colorRed.Println("App IDD required")
			continue
		}
		break
	}
	fmt.Println("Enter your crypto holdings, the currency name and the holding size. Blank currency when done.")
	for {
		position := model.Position{}
		position.Currency = crypto.Currency(strings.ToUpper(readString("  Currency: ")))
		if position.Currency == "" {
			break
		}
		holding, err := decimal.NewFromString(readString("   Holding: "))
		if err != nil || holding.Equal(decimal.Zero) {
			_, _ = colorRed.Println("Requires non zero holding.")
			continue
		}
		position.Holding = holding
		conf.Portfolio.Positions = append(conf.Portfolio.Positions, position)
	}

	fmt.Println("Read:")
	for _, position := range conf.Portfolio.Positions {
		fmt.Println("  ", position)
	}
	if readString("Is this correct (y/N)? ") == "y" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%s/%s", homeDir, model.ConfFile)
		err = model.WriteConfig(conf, fileName)
		if err != nil {
			panic(err)
		}
	}
}

func readString(prompt string) string {
	var str string
	fmt.Print(prompt)
	c, err := fmt.Scanf("%s", &str)
	if err != nil || c != 1 {
		return ""
	}
	return str
}
