package cmd

import (
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/marioharper/commuter/cmd/utils"
	"github.com/marioharper/commuter/directions"
	"github.com/spf13/cobra"
)

var Locations []directions.Location
var ConfigFile string
var cfgFile string
var from string
var to string
var numResults int
var interval int

var RootCmd = &cobra.Command{
	Use:   "commuter",
	Short: "Tool to get travel time",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// get config file location
		usr, err := user.Current()
		utils.Check(err)
		ConfigFile = fmt.Sprintf("%s/commuter-config.json", usr.HomeDir)

		if cmd.Use != "init" {
			if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
				fmt.Println("Please initialize Commuter by using the 'commuter init' command")
				os.Exit(-1)
			}

			Locations = (utils.GetLocations(ConfigFile))
		}

	},
	Run: func(cmd *cobra.Command, args []string) {

		from := Locations[utils.GetLocationByName(Locations, from)]
		to := Locations[utils.GetLocationByName(Locations, to)]
		currTime := time.Now().Unix()
		minute := 60
		interval := int64(interval * minute)
		var traveTime int64 // time leaving
		var info directions.CommuteInfo
		var shortest int = 0
		var optimalTime string

		commute := directions.Commute{
			From: from,
			To:   to,
		}

		fmt.Printf("\nCommute from %s to %s\n", commute.From.Name, commute.To.Name)
		for i := 0; i < numResults; i++ {
			var printTime string
			traveTime = currTime + (interval * int64(i))
			info = commute.GetInfo(traveTime)
			hr, min, sec := time.Unix(traveTime, 0).Clock()
			amPm := "AM"

			if hr > 12 {
				hr -= 12
				amPm = "PM"
			} else if hr == 0 {
				hr = 12
			}

			//Get best travel time out of printed results
			if shortest == 0 || int(info.TotalDuration) < shortest {
				shortest = int(info.TotalDuration)
				optimalTime = fmt.Sprintf("%d:%02d:%02d %s", hr, min, sec, amPm)
			}

			//Print results
			if i == 0 {
				printTime = "Now"
			} else {
				printTime = fmt.Sprintf("%d:%02d:%02d %s", hr, min, sec, amPm)
			}
			fmt.Printf("\n %s: %d minutes \n", printTime, int(info.TotalDuration))

			//Print optimal time for commute based on results
			if i == numResults-1 {
				fmt.Printf("\nOptimal time for commute: %s at %d minutes \n\n", optimalTime, shortest)
			}
		}

	},
}

func Execute() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}

func init() {

	RootCmd.PersistentFlags().StringVarP(&from, "from", "f", "work", "Starting location name")
	RootCmd.PersistentFlags().StringVarP(&to, "to", "t", "home", "Destination location name")
	RootCmd.PersistentFlags().IntVarP(&numResults, "number", "n", 5, "How many commute times do you want?")
	RootCmd.PersistentFlags().IntVarP(&interval, "interval", "i", 15, "How many minutes between each commute prediction?")

}
