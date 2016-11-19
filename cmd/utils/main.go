package utils

import "fmt"

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
