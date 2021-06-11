package commands

import (
	"fmt"
	model2 "github.com/nwillc/cryptoport/model"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(setupCmd)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup your portfolio configuration.",
	Long:  "Setup your portfolio configuration and save to " + model2.ConfFile,
	Run:   setup,
}

func setup(_ *cobra.Command, _ []string) {
	conf := model2.Config{
		AppID:     "",
		Portfolio: model2.Portfolio{},
	}
	fmt.Print("Enter Nomics API_ID: ")
	fmt.Scanf("%s", &conf.AppID)
	fmt.Println("Enter your crypto holdings, the currency name and the holding size. Blank currency when done.")
	for {
		position := model2.Position{}
		fmt.Print("  Currency: ")
		fmt.Scanf("%s", &position.Currency)
		if position.Currency == "" {
			break
		}
		fmt.Print("   Holding: ")
		fmt.Scanf("%f", &position.Holding)
		conf.Portfolio.Positions = append(conf.Portfolio.Positions, position)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("%s/%s", homeDir, model2.ConfFile)
	err = model2.WriteConfig(conf, fileName)
	if err != nil {
		panic(err)
	}
}
