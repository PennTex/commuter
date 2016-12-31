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

func FormatDateInput(dateInput string) int64 {
	//Current time
	currentYear, m, _ := time.Now().Date()
	currentMonth := int(m)

	//Start time
	s := strings.Split(dateInput, ":")
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

func to12Hr() {

}
