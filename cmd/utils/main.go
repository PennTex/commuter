package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Check(e error) {
	if e != nil {
		panic(e)
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

func FormatDateInput(dateInput string) (int64, error) {
	currentYear, m, _ := time.Now().Date()
	currentMonth := int(m)
	correctFormat := "MMDD:HHMM"

	//Start time
	s := strings.Split(dateInput, ":")
	if len(s) < 2 {
		return 0, fmt.Errorf("Invalid date format supplied. Expected '%s'.", correctFormat)
	}

	// Get/Validate start date and time format
	startDate, startTime := s[0], s[1]
	if len(startDate) != 4 {
		return 0, fmt.Errorf("Invalid date supplied. Date must be in MMDD format. Expected '%s'.", correctFormat)
	}
	isPM := strings.Contains(startTime, "PM")
	isAM := strings.Contains(startTime, "AM")

	if (isAM || isPM) && len(startTime) != 6 {
		return 0, fmt.Errorf("Invalid time supplied. Time must be in HHMM format. Expected '%s[AM|PM]'.", correctFormat)
	} else if (!isAM && !isPM) && len(startTime) != 4 {
		return 0, fmt.Errorf("Invalid time supplied. Time must be in HHMM format. Expected '%s'.", correctFormat)
	}

	// Get/Validate start month
	startMonth, _ := strconv.Atoi(startDate[0:2])
	if startMonth < 1 || startMonth > 12 {
		return 0, fmt.Errorf("Invalid month supplied. Month must be between 1-12. Expected '%s'.", correctFormat)
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
		return 0, fmt.Errorf("Invalid day supplied. Day must be between 1-31. Expected '%s'.", correctFormat)
	}

	// Get/Validate start time
	startHour, _ := strconv.Atoi(startTime[0:2])
	startMinute, _ := strconv.Atoi(startTime[2:4])
	if isPM || isAM {
		if startHour < 1 || startHour > 12 {
			return 0, fmt.Errorf("Invalid time supplied. Expected 12 hr time format. Expected '%s[AM|PM]'.", correctFormat)
		}
		if isPM {
			startHour += 12
		}
	} else {
		if startHour < 1 || startHour > 24 {
			return 0, fmt.Errorf("Invalid time supplied. Expected 24 hr time format. Expected %s.", correctFormat)
		}
	}

	return time.Date(startYear, time.Month(startMonth), startDay, startHour, startMinute, 0, 0, time.Local).Unix(), nil
}
