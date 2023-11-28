package log

import (
	"fmt"

	"github.com/fatih/color"
)

func Info(format string, a ...interface{}) {
	_, _ = fmt.Printf(format+"\n", a...)
}

func Green(format string, a ...interface{}) {
	color.Green(format, a...)
}

func Fatal(format string, a ...interface{}) {
	panic(fmt.Sprintf(format, a...))
}

func Warn(format string, a ...interface{}) {
	color.Yellow(format, a...)
}
