package cmd

import (
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/PennTex/commuter/cmd/config"
	"github.com/PennTex/commuter/cmd/utils"
	"github.com/PennTex/commuter/directions"
	"github.com/PennTex/commuter/weather"
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

func formatTime(unixTimeStamp int64) string {
	amPm := "AM"
	hr, min, sec := time.Unix(unixTimeStamp, 0).Clock()

	if hr > 12 {
		hr -= 12
		amPm = "PM"
	} else if hr == 0 {
		hr = 12
	}

	return fmt.Sprintf("%d:%02d:%02d %s", hr, min, sec, amPm)
}

func createCommute(fromLocationName string, toLocationName string, start string) directions.Commute {
	from, err := Config.GetLocationByName(fromLocationName)
	utils.ProcessError(err, "Getting location by name")
	to, err := Config.GetLocationByName(toLocationName)
	utils.ProcessError(err, "Getting location by name")

	if start != "" {
		if len(start) > 6 {
			commuteTime, err = utils.FormatDateTimeInput(start)
		} else {
			commuteTime, err = utils.FormatTimeInput(start)
		}
		utils.ProcessError(err, "Formatting time")
	} else {
		commuteTime = time.Now().Unix() // default
	}

	return directions.NewCommute(from, to, commuteTime)
}

func getPossibleCommutes(commute directions.Commute, minuteInterval int) []directions.Commute {
	var allCommutes []directions.Commute
	interval := int64(minuteInterval * 60)

	for i := 0; i < numResults; i++ {
		travelTime := commuteTime + (interval * int64(i))
		newCommute := directions.NewCommute(commute.From, commute.To, travelTime)

		allCommutes = append(allCommutes, newCommute)
	}

	return allCommutes
}

func getShortestCommute(commutes []directions.Commute) directions.Commute {
	var shortest directions.Commute

	for _, commute := range commutes {
		if (shortest == directions.Commute{}) || (commute.TotalDuration < shortest.TotalDuration) {
			shortest = commute
		}
	}

	return shortest
}

func printCommuteTimes(commutes []directions.Commute) {
	var printTime string

	for i, commute := range commutes {
		if i == 0 && start == "" {
			printTime = "Now"
		} else {
			printTime = formatTime(commute.Time)
		}

		fmt.Printf("\n %s: %d minutes \n", printTime, int(commute.TotalDuration))
	}
}

func printWeatherInfo(commute directions.Commute) {
	//Print Weather Info
	commuteWeather := weather.GetInfo(int(commute.Time), commute.Lat, commute.Lng)
	fmt.Printf("Summary: %v \nTemperature: %v° F \nWind Speed: %v MPH \nChance Of Rain: %v%% \n\n", commuteWeather.Summary, commuteWeather.Temp, commuteWeather.Wind, commuteWeather.PrecipProbability)
}

var RootCmd = &cobra.Command{
	Use:   "commuter",
	Short: "Tool to get travel time",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// get config file location
		usr, err := user.Current()
		utils.ProcessError(err, "")
		configFile = fmt.Sprintf("%s/commuter-config.json", usr.HomeDir)

		if cmd.Use != "init" {
			if fStat, err := os.Stat(configFile); os.IsNotExist(err) || fStat.Size() == 0 {
				fmt.Println("Please initialize Commuter by using the 'commuter init' command")
				return
			}
		}

		Config = config.New(configFile)
		commute = createCommute(from, to, start)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\nCommute from %s to %s\n", commute.From.Name, commute.To.Name)
		allCommutes := getPossibleCommutes(commute, interval)
		printCommuteTimes(allCommutes)

		shortestCommute := getShortestCommute(allCommutes)
		fmt.Printf("\nOptimal time for commute: %s at %d minutes \n", formatTime(shortestCommute.Time), int(shortestCommute.TotalDuration))

		fmt.Println("\nWeather for your commute:")
		printWeatherInfo(shortestCommute)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		utils.ProcessError(err, "")
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&from, "from", "f", "work", "Starting location name")
	RootCmd.PersistentFlags().StringVarP(&to, "to", "t", "home", "Destination location name")
	RootCmd.PersistentFlags().StringVarP(&start, "start", "s", "", "Starting commute time (1219:1330 or 1219:0130PM = December 19th at 1:30PM)")
	RootCmd.Flags().IntVarP(&numResults, "number", "n", 5, "How many commute times do you want?")
	RootCmd.Flags().IntVarP(&interval, "interval", "i", 15, "How many minutes between each commute prediction?")
}
