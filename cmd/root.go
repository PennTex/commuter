package cmd

import (
	"fmt"
	"os"

	"github.com/marioharper/commuter/directions"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var From string
var To string
var Locations []directions.Location

var RootCmd = &cobra.Command{
	Use:   "commuter",
	Short: "Tool to get travel time",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		home := directions.Location{
			Name:    "home",
			Address: "4424 Gaines Ranch Loop, Austin, TX",
		}

		work := directions.Location{
			Name:    "work",
			Address: "1835 Kramer Ln, Austin, TX 78758",
		}

		Locations = append(Locations, home)
		Locations = append(Locations, work)

	},
	Run: func(cmd *cobra.Command, args []string) {

		from := Locations[getLocationByName(Locations, From)]
		to := Locations[getLocationByName(Locations, To)]

		_, dur := directions.TotalDistDur(from, to)
		fmt.Printf("\n Commute info from %s to %s\n", from.Name, to.Name)
		fmt.Printf("\nDuration: %f minutes \n", dur)
	},
}

func getLocationByName(locations []directions.Location, name string) int {

	for idx, location := range locations {
		if location.Name == name {
			return idx
		}
	}

	return -1
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&From, "from", "f", "work", "Environment to manipulate")
	RootCmd.PersistentFlags().StringVarP(&To, "to", "t", "home", "Environment to manipulate")

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.commuter-cli.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".commuter") // name of config file (without extension)
	viper.AddConfigPath("$HOME")     // adding home directory as first search path
	viper.AutomaticEnv()             // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
