package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PennTex/commuter/directions"
)

func ProcessError(e error, errorString string) {
	if e != nil {
		if errorString != "" {
			fmt.Printf("ERROR - %s: %s \n", errorString, e.Error())
			os.Exit(-1)
		} else {
			panic(e)
		}
	}
}

type Logger struct {
	Logging bool
}

func (l *Logger) Log(output string) {
	if l.Logging {
		fmt.Println(output)
	}
}

func FormatDateTimeInput(dateInput string) (int64, error) {
	currentYear, m, _ := time.Now().Date()
	currentMonth := int(m)
	correctFormat := "MMDD:HHMM"

	//Start time
	s := strings.Split(dateInput, ":")
	if len(s) < 2 {
		return 0, fmt.Errorf("Invalid date/time format supplied. Expected '%s'. \n", correctFormat)
	}

	// Get/Validate start date and time format
	startDate, startTime := s[0], s[1]
	if len(startDate) != 4 {
		return 0, fmt.Errorf("Invalid date format. Date must be in MMDD format. Expected '%s'. \n", correctFormat)
	}
	isPM := strings.Contains(startTime, "PM")
	isAM := strings.Contains(startTime, "AM")

	if (isAM || isPM) && len(startTime) != 6 {
		return 0, fmt.Errorf("Invalid time format. Time must be in HHMM format. Expected '%s[AM|PM]'. \n", correctFormat)
	} else if (!isAM && !isPM) && len(startTime) != 4 {
		return 0, fmt.Errorf("Invalid time format. Time must be in HHMM format. Expected '%s'. \n", correctFormat)
	}

	// Get/Validate start month
	startMonth, _ := strconv.Atoi(startDate[0:2])
	if startMonth < 1 || startMonth > 12 {
		return 0, fmt.Errorf("Invalid month supplied. Month must be between 1-12. Expected '%s'. \n", correctFormat)
	}

	var startYear int
	if currentMonth > startMonth {
		startYear = currentYear + 1
	} else {
		startYear = currentYear
	}

	// Get/Validate start day
	startDay, _ := strconv.Atoi(startDate[2:4])
	if startDay < 1 || startDay > 31 {
		return 0, fmt.Errorf("Invalid day supplied. Day must be between 1-31. Expected '%s'. \n", correctFormat)
	}

	// Get/Validate start time
	startHour, _ := strconv.Atoi(startTime[0:2])
	startMinute, _ := strconv.Atoi(startTime[2:4])
	if isPM || isAM {
		if startHour < 1 || startHour > 12 {
			return 0, fmt.Errorf("Invalid time supplied. Expected 12 hr time format. Expected '%s[AM|PM]'. \n", correctFormat)
		}
		if isPM {
			startHour += 12
		}
	} else {
		if startHour < 1 || startHour > 24 {
			return 0, fmt.Errorf("Invalid time supplied. Expected 24 hr time format. Expected %s. \n", correctFormat)
		}
	}

	return time.Date(startYear, time.Month(startMonth), startDay, startHour, startMinute, 0, 0, time.Local).Unix(), nil
}

func FormatTimeInput(timeInput string) (int64, error) {
	currentYear, m, currentDay := time.Now().Date()
	currentMonth := int(m)
	correctFormat := "HHMM"
	isPM := strings.Contains(timeInput, "PM")
	isAM := strings.Contains(timeInput, "AM")

	//Start time
	if len(timeInput) != 4 && len(timeInput) != 6 {
		return 0, fmt.Errorf("Invalid time supplied. Expected 12 hr time format. Expected '%s[AM|PM]'!!! \n", correctFormat)
	}

	//Get current hour/minute
	currentHour := time.Now().Hour()
	currentMinute := time.Now().Minute()

	//Get/validate start time
	startHour, _ := strconv.Atoi(timeInput[0:2])
	startMinute, _ := strconv.Atoi(timeInput[2:4])
	if startHour < currentHour || (startHour == currentHour && startMinute < currentMinute) {
		return 0, fmt.Errorf("Invalid time supplied. Time cannot be in the past.\n")
	}
	if isPM || isAM {
		if startHour < 1 || startHour > 12 {
			return 0, fmt.Errorf("Invalid time supplied. Expected 12 hr time format. Expected '%s[AM|PM]'. \n", correctFormat)
		}
		if isPM {
			startHour += 12
		}
	} else {
		if startHour < 1 || startHour > 24 {
			return 0, fmt.Errorf("Invalid time supplied. Expected 24 hr time format. Expected %s. \n", correctFormat)
		}
	}

	return time.Date(currentYear, time.Month(currentMonth), currentDay, startHour, startMinute, 0, 0, time.Local).Unix(), nil

}

func GetLocationNameFromUser(reader *bufio.Reader) string {
	name := ""

	for name == "" {
		fmt.Print("Enter location name: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			fmt.Println("Please supply a value for the location's name.")
		}
	}

	return name
}

func GetLocationAddressFromUser(addressName string, addressValidator directions.AddressValidator, reader *bufio.Reader) string {
	address := ""

	for address == "" {
		fmt.Printf("Enter %s address: ", addressName)
		address, _ = reader.ReadString('\n')
		address = strings.TrimSpace(address)

		validAddress, err := addressValidator.IsValidAddress(address)
		ProcessError(err, "Validating address.")

		if !validAddress {
			address = ""
		}

		if address == "" {
			fmt.Printf("Please supply a valid address for '%s'. \n", addressName)
		}
	}

	return address
}
