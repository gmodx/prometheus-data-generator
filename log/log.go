package log

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func Info(format string, a ...interface{}) {
	_, _ = fmt.Printf(time.Now().Format(time.StampMilli)+" "+format+"\n", a...)
}

func Green(format string, a ...interface{}) {
	color.Green(time.Now().Format(time.StampMilli)+" "+format, a...)
}

func Fatal(format string, a ...interface{}) {
	panic(fmt.Sprintf(time.Now().Format(time.StampMilli)+" "+format, a...))
}

func Warn(format string, a ...interface{}) {
	color.Yellow(time.Now().Format(time.StampMilli)+" "+format, a...)
}
