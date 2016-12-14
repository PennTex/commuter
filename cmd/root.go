package cmd

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
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
var from string
var to string
var numResults int
var interval int
var start string

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

	},
	Run: func(cmd *cobra.Command, args []string) {

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

		minute := 60
		interval := int64(interval * minute)
		var traveTime int64 // time leaving
		var info directions.CommuteInfo
		var shortest int = 0
		var optimalTime string
		var bestHour int
		amPm := "AM"
		commuteTime := time.Now().Unix()

		if start != "" {
			commuteTime = _formatDateInput()
		}

		commute := directions.Commute{
			From: from,
			To:   to,
		}

		fmt.Printf("\nCommute from %s to %s\n", commute.From.Name, commute.To.Name)
		for i := 0; i < numResults; i++ {
			var printTime string
			traveTime = commuteTime + (interval * int64(i))
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
		var commuteWeather = weather.GetInfo(bestHour, amPm, info.Lat, info.Lng)
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

func _formatDateInput() int64 {
	//Current time
	currentYear, m, _ := time.Now().Date()
	currentMonth := int(m)

	//Start time
	s := strings.Split(start, ":")
	startDate, startTime := s[0], s[1]
	startMonth, _ := strconv.Atoi(startDate[0:2])
	var startYear int
	if currentMonth > startMonth {
		startYear = currentYear + 1
	} else {
		startYear = currentYear
	}
	startDay, _ := strconv.Atoi(startDate[2:4])
	startHour, _ := strconv.Atoi(startTime[0:2])
	startMinute, _ := strconv.Atoi(startTime[2:4])
	if strings.Contains(startTime, "PM") {
		startHour += 12
	}

	return time.Date(startYear, time.Month(startMonth), startDay, startHour, startMinute, 0, 0, time.Local).Unix()
}
