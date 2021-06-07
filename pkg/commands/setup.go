package commands

import (
	"fmt"
	"github.com/nwillc/cryptoport/pkg/model"
	"github.com/spf13/cobra"
	"os"
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
	fmt.Print("Enter Nomics API_ID: ")
	fmt.Scanf("%s", &conf.AppID)
	fmt.Println("Enter your crypto holdings, the currency name and the holding size. Blank currency when done.")
	for {
		position := model.Position{}
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
	fileName := fmt.Sprintf("%s/%s", homeDir, model.ConfFile)
	err = model.WriteConfig(conf, fileName)
	if err != nil {
		panic(err)
	}
}
