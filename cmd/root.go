package cmd

import (
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/marioharper/commuter/cmd/config"
	"github.com/marioharper/commuter/cmd/utils"
	"github.com/marioharper/commuter/directions"
	"github.com/marioharper/commuter/weather"
	"github.com/spf13/cobra"
)

var Config config.ConfigManager
var Logger = utils.Logger{Logging: false}
var configFile string
var commute directions.Commute

var from string
var to string
var numResults int
var interval int
var start string
var commuteTime int64

var RootCmd = &cobra.Command{
	Use:   "commuter",
	Short: "Tool to get travel time",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// get config file location
		usr, err := user.Current()
		utils.Check(err)
		configFile = fmt.Sprintf("%s/commuter-config.json", usr.HomeDir)

		if cmd.Use != "init" {
			if fStat, err := os.Stat(configFile); os.IsNotExist(err) || fStat.Size() == 0 {
				fmt.Println("Please initialize Commuter by using the 'commuter init' command")
				os.Exit(-1)
			}
		}

		Config = config.New(configFile)

		// setup commute object
		from, err := Config.GetLocationByName(from)
		if err != nil {
			fmt.Printf(err.Error())
			panic(err)
		}
		to, err := Config.GetLocationByName(to)
		if err != nil {
			fmt.Printf(err.Error())
			panic(err)
		}

		commuteTime = time.Now().Unix() // default

		if start != "" {
			commuteTime = utils.FormatDateInput(start)
		}

		commute = directions.Commute{
			From: from,
			To:   to,
			Time: commuteTime,
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		minute := 60
		interval := int64(interval * minute)
		var info directions.CommuteInfo
		var shortest int = 0
		var optimalTime string
		var bestHour int
		amPm := "AM"

		fmt.Printf("\nCommute from %s to %s\n", commute.From.Name, commute.To.Name)
		for i := 0; i < numResults; i++ {
			var printTime string
			traveTime := commuteTime + (interval * int64(i))
			info = commute.GetInfo(traveTime)
			hr, min, sec := time.Unix(traveTime, 0).Clock()

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
				bestHour = hr
			}

			//Print results
			if i == 0 && start == "" {
				printTime = "Now"
			} else {
				printTime = fmt.Sprintf("%d:%02d:%02d %s", hr, min, sec, amPm)
			}
			fmt.Printf("\n %s: %d minutes \n", printTime, int(info.TotalDuration))

			//Print optimal time for commute based on results
			if i == numResults-1 {
				fmt.Printf("\nOptimal time for commute: %s at %d minutes \n", optimalTime, shortest)
			}
		}

		//Print Weather Info
		commuteWeather := weather.GetInfo(bestHour, amPm, info.Lat, info.Lng)
		fmt.Println("\nWeather for your commute:")
		fmt.Printf("Summary: %v \nTemperature: %v° F \nWind Speed: %v MPH \nChance Of Rain: %v%% \n\n", commuteWeather.Summary, commuteWeather.Temp, commuteWeather.Wind, commuteWeather.PrecipProbability)
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
	RootCmd.PersistentFlags().StringVarP(&start, "start", "s", "", "Starting commute time (1219:1330 or 1219:0130PM = December 19th at 1:30PM)")
	RootCmd.Flags().IntVarP(&numResults, "number", "n", 5, "How many commute times do you want?")
	RootCmd.Flags().IntVarP(&interval, "interval", "i", 15, "How many minutes between each commute prediction?")

}
